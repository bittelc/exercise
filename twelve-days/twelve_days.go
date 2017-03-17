// Package twelve returns a single verse for any of the twelve days of Christmas
package twelve

import ()

const testVersion = 1

// Verse returns a single verse corresponding to the i'th numbered verse
func Verse(i int) string {
	return "oh ya"
}

// Song returns the entire song, line seperating each verse
func Song() string {
	// Could potentially use a recurse loop here. Eg. string += Song(i+1)
	return "string"
}
