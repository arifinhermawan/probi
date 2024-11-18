package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/reminder"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
	"github.com/jmoiron/sqlx"
)

type PSQL struct {
	Reminder *reminder.Repository
	User     *user.Repository
}

func NewPSQL(lib *lib.Lib, psql *sqlx.DB) *PSQL {
	return &PSQL{
		Reminder: reminder.NewRepository(lib, psql),
		User:     user.NewRepository(lib, psql),
	}
}
