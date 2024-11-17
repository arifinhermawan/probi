package main

import (
	"context"
	"os"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/blib/tracer"
	app "github.com/arifinhermawan/probi/internal/app/cron"
	"github.com/arifinhermawan/probi/internal/app/utils"
	"github.com/arifinhermawan/probi/internal/lib/constants"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
)

const (
	filePath = "etc/logs/application.log"
)

func main() {
	ctx := internalContext.DefaultContext()
	ctx = context.WithValue(ctx, log.ContextKey("service_type"), constants.ServiceNameCRON)

	nrApp, err := utils.InitNewRelicConn()
	if err != nil {
		log.Fatal(ctx, nil, err, "Failed to init new relic")
	}

	cleanUp(ctx)
	tracer.InitTracer(nrApp)
	log.Init(filePath)
	app.NewCRONApplication(ctx, nrApp)
}

func cleanUp(ctx context.Context) {
	if err := os.Truncate(filePath, 0); err != nil {
		log.Warn(ctx, nil, nil, "Failed to open log file")
	}
}
