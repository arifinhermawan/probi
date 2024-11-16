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
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

type UpdateReminderReq struct {
	ID        int64     `db:"id"`
	Frequency string    `db:"frequency"`
	Interval  int       `db:"interval"`
	EndDate   time.Time `db:"end_date"`
	DueDate   time.Time `db:"due_date"`
	UpdatedAt time.Time `db:"updated_at"`
}
