package model

type TTSReq struct {
	Text string `json:"text"`
}

type TTSResp struct {
	Text      string `json:"text"`
	AudioData []byte `json:"audio_data"`
}
