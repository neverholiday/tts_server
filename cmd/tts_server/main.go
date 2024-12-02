// Command quickstart generates an audio file with the content "Hello, World!".
package main

import (
	"context"
	"tts_server/cmd/tts_server/app"
	"tts_server/cmd/tts_server/repositories"
	"tts_server/thirdparty/googletts"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"github.com/labstack/echo"
)

var (
	googleTTSClient *googletts.GoogleTTSClient
	googleTTSRepo   *repositories.GoogleTextToSpeechRepo
	chatGPTTTSRepo  *repositories.ChatGPTTextToSpeechRepo
	appHandler      *app.App
)

func init() {
	// Instantiates a client.
	ctx := context.Background()

	// google text to speech
	client, err := texttospeech.NewClient(ctx)
	if err != nil {
		panic(err)
	}
	googleTTSClient = googletts.NewGoogleTTSClient(client)

	// new repositories
	googleTTSRepo = repositories.NewGoogleTextToSpeechRepo(googleTTSClient.Client)
	chatGPTTTSRepo = repositories.NewChatGPTTextToSpeechRepo()

	// new app
	appHandler = app.NewApp(googleTTSRepo, chatGPTTTSRepo)
}

func main() {

	defer googleTTSClient.Close()

	e := echo.New()
	e.POST("/synthesize/google", appHandler.GoogleSynthesizeAudio)
	e.Logger.Fatal(e.Start(":8080"))

}
