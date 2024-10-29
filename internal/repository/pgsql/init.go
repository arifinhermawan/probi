package pgsql

import (
	"context"

	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
)

type userProvider interface {
	CreateUserInDB(ctx context.Context, req user.CreateUserReq) (int64, error)
}

type DBRepo struct {
	user userProvider
}

func NewDBRepository(user userProvider) *DBRepo {
	return &DBRepo{
		user: user,
	}
}
