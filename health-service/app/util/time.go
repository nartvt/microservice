package util

import "time"

const (
	FormatTime = "2006-01-02T15:04:05.000Z"
)

func FormatDateTime(date time.Time) string {
	return date.Format(FormatTime)
}
