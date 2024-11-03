package main

import (
	"os"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app"
	"github.com/arifinhermawan/probi/internal/app/utils"
	internalContext "github.com/arifinhermawan/probi/internal/lib/context"
)

const (
	filePath = "etc/logs/application.log"
)

func main() {
	nrApp, err := utils.InitNewRelicConn()
	if err != nil {
		log.Fatal(internalContext.DefaultContext(), nil, err, "Failed to init new relic")
	}

	// clean up log file so it doesn't
	// get bloated
	cleanUp()
	log.Init(filePath)
	app.NewApplication(nrApp)
}

func cleanUp() {
	if err := os.Truncate(filePath, 0); err != nil {
		log.Warn(internalContext.DefaultContext(), nil, nil, "Failed to open log file")
	}
}
