package gigasecond

import "time"

func AddGigasecond(t time.Time) time.Time {
	t = t.Add(time.Duration(1000000000)*time.Second)
	return t
}
