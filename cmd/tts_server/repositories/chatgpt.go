package repositories

import (
	"context"
	"tts_server/cmd/tts_server/model"
)

type ChatGPTTextToSpeechRepo struct {
}

func NewChatGPTTextToSpeechRepo() *ChatGPTTextToSpeechRepo {
	return &ChatGPTTextToSpeechRepo{}
}

func (r *ChatGPTTextToSpeechRepo) SynthesizeAudio(ctx context.Context, text string) (*model.TTSAudio, error) {
	var resp model.TTSAudio
	return &resp, nil
}
