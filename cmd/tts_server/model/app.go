package model

type TTSReq struct {
	Text string `json:"text"`
}

type TTSResp struct {
	Text      string `json:"text"`
	AudioData []byte `json:"audio_data"`
}

type OpenAITTSReq struct {
	TTSReq
	SpeechVoice string `json:"speech_voice"`
	Model       string `json:"tts_model"`
}

type GoogleTTSReq struct {
	TTSReq
}
