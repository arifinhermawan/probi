package reminder

var (
	queryCreateReminderInDB = `
		INSERT INTO reminder
		(
			user_id,
			title,
			frequency,
			"interval",
			start_date,
			end_date,
			due_date,
			created_at,
			updated_at
		)
		VALUES(
			:user_id,
			:title,
			:frequency,
			:interval,
			:start_date,
			:end_date,
			:due_date,
			:created_at,
			:updated_at
		)
	`

	queryGetActiveReminderByUserIDFromDB = `
		SELECT 
			id,
			title,
			frequency,
			"interval",
			due_date
		FROM 
			reminder
		WHERE
			user_id = $1
		AND
			deleted_at IS NULL
	`

	queryUpdateReminderInDB = `
		UPDATE
			reminder
		SET
			frequency = :frequency,
			"interval" = :interval,
			due_date = :due_date,
			end_date = :end_date,
			updated_at = :updated_at
		WHERE
			id = :id
	`
)
