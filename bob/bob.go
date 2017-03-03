// Package bob defines a lame teenager who responds like a teenager does
package bob

import "regexp"
import "fmt"

const testVersion = 2

// Hey accepts a statement to Bob and responds with a teenager-response
func Hey(str string) string {

	isAllCaps := regexp.MustCompile("^[A-Z, ,\\!,\\?]{1,}$")
	isNumsAndCaps := regexp.MustCompile("(^[0-9,\\, ]{1,})+([A-Z,\\!]{1,}$)")
	isSpCharsAndCaps := regexp.MustCompile("(^[ -`]*)+([A-Z]{1,})+([ -`]*$)")
	isLatinSupplement := regexp.MustCompile("^[A-Z,!,Ä,Ü]{1,}$")
	isQuestion := regexp.MustCompile("^.*\\?{1,} *$")
	isBlank := regexp.MustCompile("^[^A-z][ ,\\t,\\n,\\r,\\v,\u00a0,\u2002]*$")

	switch {
	case isBlank.MatchString(str),
		len(str) == 0:
		return "Fine. Be that way!"
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
		fmt.Println(len(str))
		return "Whatever."
	}

}
