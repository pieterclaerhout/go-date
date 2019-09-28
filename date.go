package yddate

import (
	"time"
)

// UnixRoundToHour returns the timestamp truncated to the hour.
func UnixRoundToHour(unixTime int64) int64 {
	t := time.Unix(unixTime, 0).UTC()
	return t.Truncate(time.Hour).UTC().Unix()
}
