package authentication

type logInResponse struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}
