package reminder

import "time"

type CreateReminderReq struct {
	UserID    int64
	Title     string
	Frequency string
	Interval  int
	StartDate time.Time
	EndDate   time.Time
}
