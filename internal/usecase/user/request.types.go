package user

type CreateUserReq struct {
	Username    string
	DisplayName string
	Email       string
	Password    string
}

type UpdateUserDetailsReq struct {
	UserID      int64
	Username    string
	DisplayName string
	Email       string
}
