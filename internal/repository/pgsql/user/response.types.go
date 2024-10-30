package user

type User struct {
	ID          int64  `db:"id"`
	Username    string `db:"username"`
	DisplayName string `db:"display_name"`
	Email       string `db:"email"`
	Password    string `db:"password"`
	Status      string `db:"status"`
}
