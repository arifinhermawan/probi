package reminder

import "time"

type CreateReminderReq struct {
	UserID    int64     `db:"user_id"`
	Title     string    `db:"title"`
	Frequency string    `db:"frequency"`
	Interval  int       `db:"interval"`
	StartDate time.Time `db:"start_date"`
	EndDate   time.Time `db:"end_date"`
	DueDate   time.Time `db:"due_date"`
}
