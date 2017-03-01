// Package bob defines a lame teenager who responds like a teenager does
package bob

import "regexp"

const testVersion = 2

// Hey accepts a statement to Bob and responds with a teenager-response
func Hey(str string) string {

	//isCaps := regexp.MustCompile("\\w*[A-Z]\\b")
	//isCaps := regexp.MustCompile("\\b[A-Z, ,\\!]{1,}")
	//isCaps := regexp.MustCompile("^[A-Z,0-9, ,\\!,\\?]*$")
	isAllCaps := regexp.MustCompile("^[A-Z, ,\\!,\\?]*$")
	isNumsAndCaps := regexp.MustCompile("(^[0-9,\\, ]{1,})+([A-Z,\\!]{1,}$)")
	isSpCharsAndCaps := regexp.MustCompile("(^[ -`]*)+([A-Z]{1,})+([ -`]*$)")

	isLatinSupplement := regexp.MustCompile("")

	isQuestion := regexp.MustCompile("\\w*\\?")

	switch {
	case isAllCaps.MatchString(str):
		return "Whoa, chill out!"
	case isNumsAndCaps.MatchString(str):
		return "Whoa, chill out!"
	case isSpCharsAndCaps.MatchString(str):
		return "Whoa, chill out!"

	case isLatinSupplement.MatchString(str):
		return "Whoa, chill out!"

	case isQuestion.MatchString(str):
		return "Sure."
	default:
		return "Whatever."
	}

}
