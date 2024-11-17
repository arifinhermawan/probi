package authentication

import "github.com/go-playground/validator/v10"

type logInReq struct {
	Input    string `json:"input"`
	Password string `json:"password" validate:"required"`
}

func (r *logInReq) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
