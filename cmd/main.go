package main

import (
	"context"
	"os"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app"
)

const (
	filePath = "etc/logs/application.log"
)

func main() {
	// clean up log file so it doesn't
	// get bloated
	cleanUp()
	log.Init(filePath)
	app.NewApplication()
}

func cleanUp() {
	if err := os.Truncate(filePath, 0); err != nil {
		log.Warn(context.Background(), nil, nil, "Failed to open log file")
	}
}
