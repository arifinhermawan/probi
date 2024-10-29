package pgsql

import (
	"context"

	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
)

func (db *DBRepo) CreateUserInDB(ctx context.Context, req user.CreateUserReq) (int64, error) {
	return db.user.CreateUserInDB(ctx, req)
}
