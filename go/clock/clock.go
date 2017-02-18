package clock

import "fmt"

const testVersion = 4

type Clock struct {
	Hour, Minute int
}

func New(hour, minute int) Clock {
	return Clock{Hour: hour, Minute: minute}
}

func (c Clock) String() string {
	hour, minute := c.Hour, c.Minute
	new_minute := minute % 60
	added_hour := int(minute / 60)
	new_hour := (added_hour + hour) % 24
	if new_minute < 0 {
		new_minute += 60
		new_hour--
	}
	if new_hour < 0 {
		new_hour += 24
	}
	return fmt.Sprintf("%02d:%02d", new_hour, new_minute)
}

func (c Clock) Add(minutes int) Clock {
	return Clock{Hour: c.Hour, Minute: c.Minute + minutes}
}
