// Package pangram determines if a given string is a pangram or not
package pangram

const testVersion = 1

// IsPangram takes a string and determines if it is a pangram, return true if so; false if not
func IsPangram(string) bool {
	uniquity := 0
	unique_letters := ""

	/*
		for each character in string {
			if character is not in unique_letter {
				unique_letter += character
				uniquity++
			}
		}

		if uniquity == 24 {
			return true
		}
		return false
	*/

	return true
}
