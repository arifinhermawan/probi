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
	rabbitmq "github.com/arifinhermawan/probi/internal/repository/rabbit-mq"
	"github.com/arifinhermawan/probi/internal/repository/redis"
)

func NewHTTPApplication(ctx context.Context) {
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
	db, err := utils.InitDBConn(ctx, lib.GetConfig().Database)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewHTTPApplication] utils.InitDBConn() got error")
		return
	}
	defer db.Close()

	// init redis connection
	redisClient, err := utils.InitRedisConn(ctx, lib.GetConfig().Redis)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewHTTPApplication] utils.InitRedisConn() got error")
		return
	}
	defer redisClient.Close()

	rabbit, err := utils.InitRMQConn(ctx, lib.GetConfig().RMQ)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.InitRMQConn() got error")
		return
	}
	defer rabbit.Close()

	publisher, err := utils.InitPublisher(ctx, rabbit)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.InitPublisher() got error")
		return
	}
	defer publisher.Close()

	// init app stack
	// repo
	psql := server.NewPSQL(lib, db)
	redisRepo := redis.NewRedisRepository(redisClient)
	rabbitMQRepo := rabbitmq.NewRMQRepo(publisher)

	// service
	svc := server.NewService(lib, psql, redisRepo, rabbitMQRepo)

	// usecase
	uc := server.NewUseCases(lib, svc)

	// handler
	handler := server.NewHandler(uc)
	router.HandleRequest(ctx, lib, handler)
}
