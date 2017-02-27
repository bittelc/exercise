// Package pangram determines if a given string is a pangram or not
package pangram

import (
	"fmt"
	//"regexp"
	"strings"
)

const testVersion = 1

// IsPangram takes a string and determines if it is a pangram, return true if so; false if not
func IsPangram(str string) bool {
	uniquity := 0
	var uniqueLetters string
	//isAscii := regexp.MustCompile("[A-Z]")

	for _, char := range strings.ToUpper(str) {
		fmt.Println(string(char))
		if !strings.Contains(uniqueLetters, string(char)) {
			// isAscii.MatchString(fmt.Sprint(char)) {
			uniqueLetters += string(char)
			uniquity++
		}
	}

	if uniquity == 24 {
		return true
	}

	fmt.Printf("Uniquity = %v\n", uniquity)
	fmt.Println("Unique letters = " + uniqueLetters)
	return false
}
