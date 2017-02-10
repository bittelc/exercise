// Clock stub file

// To use the right term, this is the package *clause*.
// You can document general stuff about the package here if you like.
package clock

import "fmt"

// The value of testVersion here must match `targetTestVersion` in the file
// clock_test.go.
const testVersion = 4

// Clock API as stub definitions.  No, it doesn't compile yet.
// More details and hints are in clock_test.go.

type Clock struct {
	Minute, Hour int
}

func New(hour, minute int) Clock {
	clock := Clock{Hour: hour, Minute: minute}
	fmt.Println(clock)
	return clock
}

func (c Clock) String() string {
	return fmt.Sprintf("%b", c)
}

func (c Clock) Add(minutes int) Clock {
	b := Clock{Hour: c.Hour, Minute: c.Minute + minutes}
	return b
}
