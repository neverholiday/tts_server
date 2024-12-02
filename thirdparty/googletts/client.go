package googletts

import (
	"fmt"

	texttospeech "cloud.google.com/go/texttospeech/apiv1"
)

type GoogleTTSClient struct {
	Client *texttospeech.Client
}

func NewGoogleTTSClient(client *texttospeech.Client) *GoogleTTSClient {
	return &GoogleTTSClient{Client: client}
}

func (c *GoogleTTSClient) Close() error {
	err := c.Client.Close()
	if err != nil {
		return fmt.Errorf("unable to close: %v", err)
	}
	return nil
}
