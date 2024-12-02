package repositories

import (
	"context"
	"fmt"
	"tts_server/cmd/tts_server/model"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"cloud.google.com/go/texttospeech/apiv1/texttospeechpb"
)

type GoogleTextToSpeechRepo struct {
	Client *texttospeech.Client
}

func NewGoogleTextToSpeechRepo(client *texttospeech.Client) *GoogleTextToSpeechRepo {
	return &GoogleTextToSpeechRepo{Client: client}
}

func (r *GoogleTextToSpeechRepo) SynthesizeAudio(ctx context.Context, text string) (*model.TTSAudio, error) {

	// Perform the text-to-speech request on the text input with the selected
	// voice parameters and audio file type.
	req := texttospeechpb.SynthesizeSpeechRequest{
		// Set the text input to be synthesized.
		Input: &texttospeechpb.SynthesisInput{
			InputSource: &texttospeechpb.SynthesisInput_Text{Text: text},
		},
		// Build the voice request, select the language code ("en-US") and the SSML
		Voice: &texttospeechpb.VoiceSelectionParams{
			LanguageCode: "th-TH",
		},
		// Select the type of audio file you want returned.
		AudioConfig: &texttospeechpb.AudioConfig{
			AudioEncoding: texttospeechpb.AudioEncoding_MP3,
		},
	}

	resp, err := r.Client.SynthesizeSpeech(ctx, &req)
	if err != nil {
		return nil, fmt.Errorf("unable to synth: %v", err)
	}

	return &model.TTSAudio{Text: text, AudioData: resp.AudioContent}, nil
}
