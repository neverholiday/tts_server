package repositories

import (
	"context"
	"fmt"
	"io"
	"tts_server/cmd/tts_server/model"

	"github.com/sashabaranov/go-openai"
)

type OpenAITextToSpeechRepo struct {
	Client *openai.Client
}

func NewOpenAITextToSpeechRepo(client *openai.Client) *OpenAITextToSpeechRepo {
	return &OpenAITextToSpeechRepo{Client: client}
}

func (r *OpenAITextToSpeechRepo) SynthesizeAudio(ctx context.Context, req model.OpenAITTSReq) (*model.TTSAudio, error) {

	resp, err := r.Client.CreateSpeech(ctx, openai.CreateSpeechRequest{
		Model: openai.SpeechModel(req.Model),
		Voice: openai.SpeechVoice(req.SpeechVoice),
		Input: req.Text,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to synth: %v", err)
	}

	// read response
	audioData, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}
	return &model.TTSAudio{Text: req.Text, AudioData: audioData}, nil
}
