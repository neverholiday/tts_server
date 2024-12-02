package app

import (
	"net/http"
	"tts_server/cmd/tts_server/model"
	"tts_server/cmd/tts_server/repositories"

	"github.com/labstack/echo"
)

type App struct {
	GoogleTTSRepo *repositories.GoogleTextToSpeechRepo
	OpenAITTSRepo *repositories.OpenAITextToSpeechRepo
}

func NewApp(googleTTS *repositories.GoogleTextToSpeechRepo,
	openAITTS *repositories.OpenAITextToSpeechRepo) *App {
	return &App{GoogleTTSRepo: googleTTS,
		OpenAITTSRepo: openAITTS}
}

func (a *App) GoogleSynthesizeAudio(c echo.Context) error {

	ctx := c.Request().Context()

	var req model.GoogleTTSReq
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	audio, err := a.GoogleTTSRepo.SynthesizeAudio(ctx, req)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	// convert to base64
	// audioDataStr := base64.StdEncoding.EncodeToString(audio.AudioData)

	resp := model.TTSResp{
		Text:      audio.Text,
		AudioData: audio.AudioData,
	}

	return c.JSON(http.StatusOK, resp)
}

func (a *App) OpenAISynthesizeAudio(c echo.Context) error {

	ctx := c.Request().Context()

	var req model.OpenAITTSReq
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	audio, err := a.OpenAITTSRepo.SynthesizeAudio(ctx, req)
	if err != nil {
		return c.NoContent(http.StatusInternalServerError)
	}

	resp := model.TTSResp{
		Text:      audio.Text,
		AudioData: audio.AudioData,
	}

	return c.JSON(http.StatusOK, resp)
}
