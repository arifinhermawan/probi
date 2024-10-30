package user

import (
	"context"
	"database/sql"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
}

type psqlProvider interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	GetContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
	Rebind(query string) string
}

type Repository struct {
	lib libProvider
	db  psqlProvider
}

func NewRepository(lib libProvider, db psqlProvider) *Repository {
	return &Repository{
		lib: lib,
		db:  db,
	}
}
