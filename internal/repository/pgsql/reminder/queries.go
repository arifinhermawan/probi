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
)
