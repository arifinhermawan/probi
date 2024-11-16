package reminder

import "time"

type Reminder struct {
	ID        int64     `db:"id"`
	Title     string    `db:"title"`
	Frequency string    `db:"frequency"`
	Interval  int       `db:"interval"`
	DueDate   time.Time `db:"due_date"`
}
