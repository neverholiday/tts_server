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

func (r *OpenAITextToSpeechRepo) SynthesizeAudio(ctx context.Context, text string) (*model.TTSAudio, error) {

	resp, err := r.Client.CreateSpeech(ctx, openai.CreateSpeechRequest{
		Model: openai.SpeechModel(openai.TTSModel1),
		Voice: openai.VoiceAlloy,
		Input: text,
	})

	if err != nil {
		return nil, fmt.Errorf("unable to synth: %v", err)
	}

	// read response
	audioData, err := io.ReadAll(resp)
	if err != nil {
		return nil, fmt.Errorf("unable to read response: %v", err)
	}
	return &model.TTSAudio{Text: text, AudioData: audioData}, nil
}
