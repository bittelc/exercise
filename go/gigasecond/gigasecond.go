// Package gigaseconds calculates the datetime after a given datetime in which 1e9 seconds have passed
package gigasecond

import (
	"time"
)

const testVersion = 4
const Gigasecond = time.Duration(1e9) * time.Second

// API function.  It uses a type from the Go standard library.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(Gigasecond)
}
