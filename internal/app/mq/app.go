package mq

import (
	"context"
	"log"

	"github.com/arifinhermawan/probi/internal/app/mq/router"
	"github.com/arifinhermawan/probi/internal/app/mq/server"
	"github.com/arifinhermawan/probi/internal/app/utils"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/lib/time"
	rabbitmq "github.com/arifinhermawan/probi/internal/repository/rabbit-mq"
	"github.com/arifinhermawan/probi/internal/repository/redis"
)

func NewMQApplication(ctx context.Context) {
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
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.InitDBConn() got error")
		return
	}
	defer db.Close()

	redisClient, err := utils.InitRedisConn(ctx, lib.GetConfig().Redis)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.InitRedisConn() got error")
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

	err = utils.DeclareExchange(lib.GetConfig().Channel, publisher)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.DeclareExchange() got error")
		return
	}

	consumer, err := utils.InitConsumer(ctx, rabbit)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.InitConsumer() got error")
		return
	}
	defer consumer.Close()

	err = utils.DeclareQueueAndBind(lib.GetConfig().Channel, consumer)
	if err != nil {
		log.Fatal(ctx, nil, err, "[NewMQApplication] utils.DeclareQueueAndBind() got error")
		return
	}

	psql := server.NewPSQL(lib, db)
	services := server.NewService(lib, psql, redis.NewRedisRepository(redisClient), rabbitmq.NewRMQRepo(publisher))
	usecases := server.NewUseCases(lib, services)
	handlers := server.NewHandler(lib, usecases, consumer)
	router.RegisterConsumer(ctx, handlers)
}
