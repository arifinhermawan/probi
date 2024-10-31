package app

import (
	"github.com/arifinhermawan/blib/log"
	route "github.com/arifinhermawan/probi/internal/app/router"
	"github.com/arifinhermawan/probi/internal/app/server"
	"github.com/arifinhermawan/probi/internal/app/utils"
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/lib/auth"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/lib/context"
	"github.com/arifinhermawan/probi/internal/lib/time"
	"github.com/arifinhermawan/probi/internal/repository/redis"
)

func NewApplication() {
	ctx := context.DefaultContext()

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
	route.HandleRequest(lib, handler)
}
