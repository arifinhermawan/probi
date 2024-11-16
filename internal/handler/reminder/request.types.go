package reminder

import (
	"github.com/go-playground/validator/v10"
)

type createReminderReq struct {
	Title     string `json:"title" validate:"required,min=3,max=100"`
	Frequency string `json:"frequency,omitempty"`
	Interval  int    `json:"interval" validate:"required"`
	StartDate string `json:"start_date" validate:"required"`
	EndDate   string `json:"end_date,omitempty"`
}

func validate(r interface{}) error {
	validate := validator.New()
	return validate.Struct(r)
}
