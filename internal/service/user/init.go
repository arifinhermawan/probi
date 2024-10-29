package user

import (
	"context"
	"time"

	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
)

type libProvider interface {
	GetTimeGMT7() time.Time
}
type dbProvider interface {
	CreateUserInDB(ctx context.Context, req user.CreateUserReq) error
}

type Service struct {
	lib libProvider
	db  dbProvider
}

func NewService(lib libProvider, db dbProvider) *Service {
	return &Service{
		lib: lib,
		db:  db,
	}
}
