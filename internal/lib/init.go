package lib

import (
	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

type configProvider interface {
	GetConfig() *configuration.AppConfig
}

type Lib struct {
	config configProvider
}

func New(config configProvider) *Lib {
	return &Lib{
		config: config,
	}
}
