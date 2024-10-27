package app

import (
	route "github.com/arifinhermawan/probi/internal/app/router"
	"github.com/arifinhermawan/probi/internal/app/server"
	"github.com/arifinhermawan/probi/internal/app/utils"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

func NewApplication() {
	// initialize infrastructure
	infra := lib.New(
		configuration.New(),
	)

	// init db connection
	_, _ = utils.InitDBConn(infra.GetConfig().Database)

	// init redis connection
	_, _ = utils.InitRedisConn(&infra.GetConfig().Redis)

	// init app stack
	// service
	svc := server.NewService()

	// usecase
	uc := server.NewUseCases(svc)

	// handler
	handler := server.NewHandler(uc)
	route.HandleRequest(handler)
}
