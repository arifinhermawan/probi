package time

import "time"

func (t *Time) GetTimeGMT7() time.Time {
	location, _ := time.LoadLocation("Asia/Jakarta")
	return time.Now().In(location)
}
