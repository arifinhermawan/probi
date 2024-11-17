package cron

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app/cron/scheduler"
	"github.com/arifinhermawan/probi/internal/app/cron/server"
	"github.com/arifinhermawan/probi/internal/app/utils"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/lib/time"
	"github.com/arifinhermawan/probi/internal/repository/redis"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewCRONApplication(ctx context.Context, app *newrelic.Application) {
	cfg := configuration.New()
	auth := auth.NewAuth(cfg)
	time := time.New()

	lib := lib.New(
		auth,
		cfg,
		time,
	)

	db, err := utils.InitDBConn(ctx, lib.GetConfig().Database)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewCRONApplication] utils.InitDBConn() got error")
		return
	}

	redisClient, err := utils.InitRedisConn(ctx, lib.GetConfig().Redis)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewCRONApplication] utils.InitRedisConn() got error")
		return
	}

	psql := server.NewPSQL(lib, db)
	services := server.NewService(lib, psql, redis.NewRedisRepository(redisClient))
	usecases := server.NewUseCases(lib, services)
	handlers := server.NewHandler(usecases)

	scheduler.RegisterScheduler(ctx, lib.GetConfig().Cron, handlers)
}
