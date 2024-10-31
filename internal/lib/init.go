package lib

import (
	"net/http"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

type authProvider interface {
	AuthMiddleware(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc
}
type configProvider interface {
	GetConfig() *configuration.AppConfig
}

type timeProvider interface {
	GetTimeGMT7() time.Time
}

type Lib struct {
	auth   authProvider
	config configProvider
	time   timeProvider
}

func New(auth authProvider, config configProvider, time timeProvider) *Lib {
	return &Lib{
		auth:   auth,
		config: config,
		time:   time,
	}
}
