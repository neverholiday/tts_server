// Command quickstart generates an audio file with the content "Hello, World!".
package main

import (
	"context"
	"encoding/base64"
	"tts_server/cmd/tts_server/app"
	"tts_server/cmd/tts_server/model"
	"tts_server/cmd/tts_server/repositories"
	"tts_server/thirdparty/googletts"
	"tts_server/thirdparty/openaitts"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/kelseyhightower/envconfig"
	"github.com/labstack/echo"
	"github.com/sashabaranov/go-openai"
	"google.golang.org/api/option"
)

var (
	envCfg          model.EnvCfg
	googleTTSClient *googletts.GoogleTTSClient
	appHandler      *app.App
)

func init() {
	// Instantiates a client.
	ctx := context.Background()

	// parse env config
	err := envconfig.Process("TTS_SERVER", &envCfg)
	if err != nil {
		panic(err)
	}

	creds, err := base64.StdEncoding.DecodeString(envCfg.GoogleTTSCredentials)
	if err != nil {
		panic(err)
	}

	// google text to speech
	gclient, err := texttospeech.NewClient(ctx, option.WithCredentialsJSON(creds))
	if err != nil {
		panic(err)
	}
	googleTTSClient = googletts.NewGoogleTTSClient(gclient)

	// openai client
	openAIClient := openai.NewClient(envCfg.OpenAIAPIKey)
	openAITTSClient := openaitts.NewOpenAITTSClient(openAIClient)

	// new repositories
	googleTTSRepo := repositories.NewGoogleTextToSpeechRepo(googleTTSClient.Client)
	openAITTSRepo := repositories.NewOpenAITextToSpeechRepo(openAITTSClient.Client)

	// new app
	appHandler = app.NewApp(googleTTSRepo, openAITTSRepo)
}

func main() {

	defer googleTTSClient.Close()

	e := echo.New()
	e.POST("/synthesize/google", appHandler.GoogleSynthesizeAudio)
	e.POST("/synthesize/openai", appHandler.OpenAISynthesizeAudio)
	e.Logger.Fatal(e.Start(":8080"))

}
