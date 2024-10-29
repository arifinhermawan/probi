package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
	"github.com/jmoiron/sqlx"
)

type PSQL struct {
	User *user.Repository
}

func NewPSQL(lib *lib.Lib, psql *sqlx.DB) *PSQL {
	return &PSQL{
		User: user.NewRepository(lib, psql),
	}
}
