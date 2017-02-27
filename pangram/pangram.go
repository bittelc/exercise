// Package pangram determines if a given string is a pangram or not
package pangram

import (
	"regexp"
	"strings"
)

const testVersion = 1

// IsPangram takes a string and determines if it is a pangram, return true if so; false if not
func IsPangram(str string) bool {
	uniquity := 0
	var uniqueLetters string
	isAscii := regexp.MustCompile("[A-Z]")

	for _, char := range strings.ToUpper(str) {
		sChar := string(char)
		if !strings.Contains(uniqueLetters, sChar) &&
			isAscii.MatchString(sChar) {
			uniqueLetters += sChar
			uniquity++
		}
	}

	if uniquity == 26 {
		return true
	}

	return false
}
