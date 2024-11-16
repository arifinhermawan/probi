package time

import "time"

func (t *Time) GetTimeGMT7() time.Time {
	return time.Now().In(LocationJakarta)
}

func (t *Time) ConvertToGMT7(input time.Time) time.Time {
	return input.In(LocationJakarta)
}
