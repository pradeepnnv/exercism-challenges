package gigasecond

import (
	"time"
)

// AddGigasecond returns the Time after adding one gigasecond to the given Time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(time.Second * 1000000000)
}
