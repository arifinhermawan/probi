package app

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app/http/router"
	"github.com/arifinhermawan/probi/internal/app/http/server"
	"github.com/arifinhermawan/probi/internal/app/http/utils"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/lib/time"
	"github.com/arifinhermawan/probi/internal/repository/redis"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewHTTPApplication(ctx context.Context, app *newrelic.Application) {
	cfg := configuration.New()
	auth := auth.NewAuth(cfg)
	time := time.New()

	// initialize lib
	lib := lib.New(
		auth,
		cfg,
		time,
	)

	// init db connection
	db, err := utils.InitDBConn(lib.GetConfig().Database)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewApplication] utils.InitDBConn() got error")
		return
	}

	// init redis connection
	redisClient, err := utils.InitRedisConn(lib.GetConfig().Redis)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewApplication] utils.InitRedisConn() got error")
		return
	}

	// init app stack
	// repo
	psql := server.NewPSQL(lib, db)
	redisRepo := redis.NewRedisRepository(redisClient)

	// service
	svc := server.NewService(lib, psql, redisRepo)

	// usecase
	uc := server.NewUseCases(svc)

	// handler
	handler := server.NewHandler(uc)
	router.HandleRequest(ctx, lib, handler)
}
