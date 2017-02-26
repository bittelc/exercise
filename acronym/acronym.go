// Package acronym converts a given string into its acronym
// Eg. "Is this a sentence" => "ITAS"
package acronym

import (
	"regexp"
	"strings"
)

const testVersion = 2

// Abbreviate takes a phrase and converts it to an acronym
func Abbreviate(str string) string {
	var acronym string
	matches := regexp.
		MustCompile("[A-Z]+[a-z]*|[a-z]+").
		FindAllString(str, -1)

	for _, word := range matches {
		acronym += strings.ToUpper(string(word[0]))
	}
	return acronym
}
