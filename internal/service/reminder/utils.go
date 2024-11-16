package reminder

import (
	"math"
	"time"

	internalTime "github.com/arifinhermawan/probi/internal/lib/time"
)

func (svc *Service) calculateDueDate(req calculateDueDateReq) time.Time {
	startDate := svc.lib.ConvertToGMT7(req.startDate)
	start := time.Date(startDate.Year(), startDate.Month(), startDate.Day(), 0, 0, 0, 0, internalTime.LocationJakarta)

	switch req.frequency {
	case "DAILY":
		return start.AddDate(0, 0, 1)
	case "WEEKLY":
		return start.AddDate(0, 0, getWeeklyDate(int(startDate.Weekday()), req.interval))
	case "MONTHLY":
		return time.Date(start.Year(), start.Month()+1, req.interval, 0, 0, 0, 0, internalTime.LocationJakarta)
	default:
		return req.endDate
	}
}

func getWeeklyDate(curr, target int) int {
	calc := target - curr
	daysToAdd := int(math.Abs(float64(calc))) % 7
	return 7 - daysToAdd
}
