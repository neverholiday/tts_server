package model

type EnvCfg struct {
	OpenAIAPIKey         string `envconfig:"OPENAI_APIKEY" required:"true"`
	GoogleTTSCredentials string `envconfig:"GOOGLE_TTS_CREDS" required:"true"`
}
