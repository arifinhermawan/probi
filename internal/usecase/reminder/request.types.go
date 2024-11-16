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

type UpdateReminderReq struct {
	ID        int64
	UserID    int64
	Frequency string
	Interval  int
	EndDate   time.Time
}
