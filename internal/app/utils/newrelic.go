package utils

import (
	"os"

	"github.com/arifinhermawan/blib/log"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func InitNewRelicConn() (*newrelic.Application, error) {
	app, err := newrelic.NewApplication(
		newrelic.ConfigAppName("probi"),
		newrelic.ConfigLicense(os.Getenv("NEW_RELIC_LICENSE_KEY")),
		newrelic.ConfigDistributedTracerEnabled(true),
	)
	if err != nil {
		log.Fatal(internalContext.DefaultContext(), nil, err, "[InitNewRelicConn] newrelic.NewApplication() got error")
	}

	return app, err
}
