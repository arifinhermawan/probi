package lib

import (
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

type configProvider interface {
	GetConfig() *configuration.AppConfig
}

type timeProvider interface {
	GetTimeGMT7() time.Time
}

type Lib struct {
	config configProvider
	time   timeProvider
}

func New(config configProvider, time timeProvider) *Lib {
	return &Lib{
		config: config,
		time:   time,
	}
}
