package reminder

import "time"

type Reminder struct {
	ID        int64
	Title     string
	Frequency string
	Interval  int
	DueDate   time.Time
}
