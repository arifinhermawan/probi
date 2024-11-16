package lib

import (
	"net/http"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

func (i *Lib) AuthMiddleware(endpointHandler func(writer http.ResponseWriter, request *http.Request)) http.HandlerFunc {
	return i.auth.AuthMiddleware(endpointHandler)
}

func (i *Lib) GetConfig() *configuration.AppConfig {
	return i.config.GetConfig()
}

func (i *Lib) GetTimeGMT7() time.Time {
	return i.time.GetTimeGMT7()
}

func (i *Lib) ConvertToGMT7(input time.Time) time.Time {
	return i.time.ConvertToGMT7(input)
}
