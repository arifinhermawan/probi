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

	queryGetUserByEmailFromDB = `
		SELECT
			id,
			username,
			display_name,
			email,
			"password"
		FROM
			"user"
		WHERE
			email = $1
	`

	queryGetUserByIDFromDB = `
		SELECT
			id,
			username,
			display_name,
			email,
			"password"
		FROM
			"user"
		WHERE
			id = $1
	`

	queryGetUserByUsernameFromDB = `
		SELECT
			id,
			username,
			display_name,
			email,
			"password"
		FROM
			"user"
		WHERE
			username = $1
	`

	queryUpdateUserInDB = `
		UPDATE
			"user"
		SET
			username = :username,
			display_name = :display_name,
			updated_at = :updated_at
		WHERE
			id = :id
	`
)
