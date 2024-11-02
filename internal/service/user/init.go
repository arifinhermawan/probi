package user

import (
	"context"
	"database/sql"
	"time"

	"github.com/arifinhermawan/probi/internal/lib/configuration"
	"github.com/arifinhermawan/probi/internal/repository/pgsql/user"
)

type libProvider interface {
	GetConfig() *configuration.AppConfig
	GetTimeGMT7() time.Time
}

type dbProvider interface {
	BeginTX(ctx context.Context, options *sql.TxOptions) (*sql.Tx, error)
	CreateUserInDB(ctx context.Context, req user.CreateUserReq) error
	GetUserByEmailFromDB(ctx context.Context, email string) (user.User, error)
	GetUserByIDFromDB(ctx context.Context, userID int64) (user.User, error)
	GetUserByUsernameFromDB(ctx context.Context, username string) (user.User, error)
	UpdateUserDetailsInDB(ctx context.Context, tx *sql.Tx, req user.UpdateUserDetailsReq) error
}

type redisProvider interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type Service struct {
	lib   libProvider
	db    dbProvider
	redis redisProvider
}

func NewService(lib libProvider, db dbProvider, redis redisProvider) *Service {
	return &Service{
		lib:   lib,
		db:    db,
		redis: redis,
	}
}
