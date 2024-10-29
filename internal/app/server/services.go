package server

import (
	"github.com/arifinhermawan/probi/internal/lib"
	"github.com/arifinhermawan/probi/internal/service/authentication"
	"github.com/arifinhermawan/probi/internal/service/user"
)

type Services struct {
	Auth *authentication.Service
	User *user.Service
}

func NewService(lib *lib.Lib, db *PSQL) *Services {
	return &Services{
		Auth: authentication.NewService(lib),
		User: user.NewService(lib, db.User),
	}
}
