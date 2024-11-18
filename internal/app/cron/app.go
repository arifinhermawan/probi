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
)

func NewCRONApplication(ctx context.Context) {
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
	log.Info(ctx, nil, nil, "[NewCRONApplication] connected to database")

	redis, err := utils.InitRedisConn(ctx, lib.GetConfig().Redis)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewCRONApplication] utils.InitRedisConn() got error")
		return
	}
	log.Info(ctx, nil, nil, "[NewCRONApplication] connected to redis")

	publisher, err := utils.InitNSQProducer(ctx, lib.GetConfig().NSQ.NSQD)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewCRONApplication] utils.InitNSQProducer() got error")
		return
	}
	log.Info(ctx, nil, nil, "[NewCRONApplication] publisher initialized")

	repo := server.NewRepository(lib, db, redis, publisher)
	services := server.NewService(lib, repo)
	usecases := server.NewUseCases(lib, services)
	handlers := server.NewHandler(usecases)

	scheduler.RegisterScheduler(ctx, lib.GetConfig().Cron, handlers)
	utils.GracefulShutDownProducer(ctx, publisher)
}
