package app

import (
	"context"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/app/http/router"
	"github.com/arifinhermawan/probi/internal/app/http/server"
	"github.com/arifinhermawan/probi/internal/app/utils"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/lib/time"
)

func NewHTTPApplication(ctx context.Context) {
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
		log.Fatal(ctx, nil, err, "[NewHTTPApplication] utils.InitDBConn() got error")
		return
	}
	log.Info(ctx, nil, nil, "[NewHTTPApplication] connected to database")

	redis, err := utils.InitRedisConn(ctx, lib.GetConfig().Redis)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewHTTPApplication] utils.InitRedisConn() got error")
		return
	}
	log.Info(ctx, nil, nil, "[NewHTTPApplication] connected to redis")

	publisher, err := utils.InitNSQProducer(ctx, lib.GetConfig().NSQ.NSQD)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewHTTPApplication] utils.InitNSQProducer() got error")
		return
	}
	log.Info(ctx, nil, nil, "[NewHTTPApplication] publisher initialized")

	repo := server.NewRepository(lib, db, redis, publisher)
	services := server.NewService(lib, repo)
	usecases := server.NewUseCases(lib, services)
	handlers := server.NewHandler(usecases)
	router.HandleRequest(ctx, lib, handlers)
	utils.GracefulShutDownProducer(ctx, publisher)
}
