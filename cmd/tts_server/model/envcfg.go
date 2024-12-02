package model

type EnvCfg struct {
	OpenAIAPIKey string `envconfig:"OPENAI_APIKEY" required:"true"`
}
