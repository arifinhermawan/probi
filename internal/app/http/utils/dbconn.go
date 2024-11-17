package utils

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"github.com/arifinhermawan/blib/log"
	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/lib/context"
)

func InitDBConn(cfg configuration.DatabaseConfig) (*sqlx.DB, error) {
	ctx := context.DefaultContext()
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DatabaseName)

	db, err := sqlx.Open(cfg.Driver, psqlInfo)
	if err != nil {
		log.Error(ctx, nil, err, "[InitDBConn] sqlx.Open() got error")
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		log.Error(ctx, nil, err, "[InitDBConn] db.Ping() got error")
		return nil, err
	}

	log.Info(ctx, nil, nil, "successfully connect to database")
	return db, nil
}
