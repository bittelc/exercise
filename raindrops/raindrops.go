// Package raindrops converts
package raindrops

import "strconv"

const (
	testVersion = 2
	Pling       = "Pling"
	Plang       = "Plang"
	Plong       = "Plong"
)

// Convert converst an integer parameter into a raindrop sound
func Convert(i int) string {
	var noise string

	if i%3 == 0 || i%5 == 0 || i%7 == 0 {
		if i%3 == 0 {
			noise += Pling
		}
		if i%5 == 0 {
			noise += Plang
		}
		if i%7 == 0 {
			noise += Plong
		}
		return noise
	}

	return strconv.Itoa(i)
}
