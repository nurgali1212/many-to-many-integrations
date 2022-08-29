package integrations

import (
	"sign_in/config"
)

type client struct {
	conf *config.Config
}

func NewClient(config *config.Config) *client {
	return &client{
		conf: config,
	}
}

type Client interface {
	SendData(err error) error
}
