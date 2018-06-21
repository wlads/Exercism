package clock

import "fmt"

// Clock represents a generic 24h time "hour:minute" e.g. 10:30, 18:05
type Clock struct {
	hour   int
	minute int
}

// New creates a new Time struct with hour and minute as parameters
func New(hour, minute int) Clock {
	for minute < 0 {
		hour--
		minute += 60
	}
	for hour < 0 {
		hour += 24
	}
	for minute >= 60 {
		hour++
		minute -= 60
	}
	for hour >= 24 {
		hour -= 24
	}
	return Clock{hour: hour, minute: minute}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add increments the minutes returning a new Clock
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute+minutes)
}

// Subtract decrements the minutes returning a new Clock
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute-minutes)
}
