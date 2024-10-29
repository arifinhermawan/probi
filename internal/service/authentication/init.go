package authentication

import "github.com/arifinhermawan/probi/internal/lib/configuration"

type libProvider interface {
	GetConfig() *configuration.AppConfig
}

type Service struct {
	lib libProvider
}

func NewService(lib libProvider) *Service {
	return &Service{
		lib: lib,
	}
}
