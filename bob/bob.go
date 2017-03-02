// Package bob defines a lame teenager who responds like a teenager does
package bob

import "regexp"
import "fmt"

const testVersion = 2

// Hey accepts a statement to Bob and responds with a teenager-response
func Hey(str string) string {

	isAllCaps := regexp.MustCompile("^[A-Z, ,\\!,\\?]*$")
	isNumsAndCaps := regexp.MustCompile("(^[0-9,\\, ]{1,})+([A-Z,\\!]{1,}$)")
	isSpCharsAndCaps := regexp.MustCompile("(^[ -`]*)+([A-Z]{1,})+([ -`]*$)")
	isLatinSupplement := regexp.MustCompile("^[A-Z,!,Ä,Ü]*$")
	isQuestion := regexp.MustCompile("^.*\\?$")
	isBlank := regexp.MustCompile("")

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
	case isBlank(str):
		return "Fine. Be that way!"
	default:
		fmt.Println(len(str))
		return "Whatever."
	}

}
