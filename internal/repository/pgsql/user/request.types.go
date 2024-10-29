package user

import "time"

type CreateUserReq struct {
	Username    string    `db:"username"`
	DisplayName string    `db:"display_name"`
	Email       string    `db:"email"`
	Password    string    `db:"password"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}
