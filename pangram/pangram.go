// Package pangram determines if a given string is a pangram or not
package pangram

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 1

// IsPangram takes a string and determines if it is a pangram, return true if so; false if not
func IsPangram(str string) bool {
	uniquity := 0
	var unique_letters string
	is_ascii := regexp.MustCompile("[A-Z]")

	for _, char := range strings.ToUpper(str) {
		if !strings.Contains(unique_letters, fmt.Sprint(char)) &&
			is_ascii.MatchString(fmt.Sprintf("%q", char)) {
			unique_letters += fmt.Sprint(char)
			uniquity++
		}
	}

	if uniquity == 24 {
		return true
	}
	fmt.Println(uniquity)
	fmt.Printf("%q", unique_letters)
	return false

}
