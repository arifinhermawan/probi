package reminder

import "time"

type Reminder struct {
	ID        int64     `json:"id"`
	Title     string    `json:"title"`
	Frequency string    `json:"frequency"`
	Interval  int       `json:"interval"`
	DueDate   time.Time `json:"due_date"`
}
