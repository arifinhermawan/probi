package configuration

import "sync"

type Configuration struct {
	config           AppConfig
	doLoadConfigOnce *sync.Once
}

func New() *Configuration {
	return &Configuration{
		doLoadConfigOnce: new(sync.Once),
	}
}
