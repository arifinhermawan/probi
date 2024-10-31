package auth

import "github.com/arifinhermawan/probi/internal/lib/configuration"

type configProvider interface {
	GetConfig() *configuration.AppConfig
}

type Auth struct {
	cfg configProvider
}

func NewAuth(cfg *configuration.Configuration) *Auth {
	return &Auth{
		cfg: cfg,
	}
}
