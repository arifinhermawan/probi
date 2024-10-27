package lib

import (
	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

func (i *Lib) GetConfig() *configuration.AppConfig {
	return i.config.GetConfig()
}
