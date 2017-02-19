// Package gigasecond calculates the datetime after a given datetime in which 1e9 seconds have passed.
package gigasecond

import (
	"time"
)

const testVersion = 4
const Gigasecond = time.Duration(1e9) * time.Second

// AddGigasecond adds a gigasecond (1 * 10^9 seconds) to any provided golang time.Time.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
