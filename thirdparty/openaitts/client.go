package openaitts

import "github.com/sashabaranov/go-openai"

type OpenAITTS struct {
	Client *openai.Client
}

func NewOpenAITTSClient(client *openai.Client) *OpenAITTS {
	return &OpenAITTS{Client: client}
}
