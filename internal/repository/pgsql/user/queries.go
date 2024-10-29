package user

var (
	queryCreateUserInDB = `
		INSERT INTO "user"(username, display_name, email, "password", created_at, updated_at)
		VALUES(
			:username,
			:display_name,
			:email,
			:password,
			:created_at,
			:updated_at
		)
	`
)
