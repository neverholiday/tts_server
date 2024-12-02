package app

import (
	"context"
	"net/http"
	"tts_server/cmd/tts_server/model"

	"github.com/labstack/echo"
)

type ITextToSpeechRepo interface {
	SynthesizeAudio(ctx context.Context, text string) (*model.TTSAudio, error)
}

type App struct {
	GoogleTTSRepo  ITextToSpeechRepo
	ChatGPTTTSRepo ITextToSpeechRepo
}

func NewApp(googleTTS ITextToSpeechRepo,
	chatgptTTS ITextToSpeechRepo) *App {
	return &App{GoogleTTSRepo: googleTTS,
		ChatGPTTTSRepo: chatgptTTS}
}

func (a *App) GoogleSynthesizeAudio(c echo.Context) error {

	ctx := c.Request().Context()

	var req model.TTSReq
	err := c.Bind(&req)
	if err != nil {
		return c.NoContent(http.StatusBadRequest)
	}

	audio, err := a.GoogleTTSRepo.SynthesizeAudio(ctx, req.Text)
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

// func (a *App) OpenAISynthesizeAudio(c echo.Context) error {

// 	ctx := c.Request().Context()

// 	var req model.TTSReq
// 	err := c.Bind(&req)
// 	if err != nil {
// 		return c.NoContent(http.StatusBadRequest)
// 	}

// 	audio, err := a.ChatGPTTTSRepo.SynthesizeAudio(ctx, req.Text)
// 	if err != nil {
// 		return c.NoContent(http.StatusInternalServerError)
// 	}

// 	return nil
// }
