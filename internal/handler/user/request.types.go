package user

import "github.com/go-playground/validator/v10"

type createUserReq struct {
	Username    string `json:"username" validate:"required,min=3,max=30"`
	DisplayName string `json:"display_name" validate:"required,min=3,max=50"`
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" validate:"required"`
}

func (r *createUserReq) validate() error {
	validate := validator.New()
	return validate.Struct(r)
}
