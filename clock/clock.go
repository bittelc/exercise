// clock is a package about clocks and allws for adding subtracting and comparing clocks
package clock

import "fmt"

const testVersion = 4

type Clock struct {
	Minute int
}

func New(hour, minute int) Clock {
	c := Clock{Minute: (60*hour + minute) % 1440}
	if c.Minute < 0 {
		c.Minute += 1440
	}
	return c
}

func (c Clock) String() string {
	hourOutput := int(c.Minute / 60)
	minuteOutput := int(c.Minute % 60)
	return fmt.Sprintf("%02d:%02d", hourOutput, minuteOutput)
}

func (c Clock) Add(minutes int) Clock {
	c.Minute = (c.Minute + minutes) % 1440
	if c.Minute < 0 {
		c.Minute += 1440
	}
	return c
}
