// Package bob defines a lame teenager who responds like a teenager does
package bob

import "regexp"

const (
	testVersion = 2
	sure        = "Sure."
	whatever    = "Whatever."
	chill       = "Whoa, chill out!"
	fine        = "Fine. Be that way!"
)

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
		return fine
	case isAllCaps.MatchString(str),
		isNumsAndCaps.MatchString(str),
		isSpCharsAndCaps.MatchString(str),
		isLatinSupplement.MatchString(str):
		return chill
	case isQuestion.MatchString(str):
		return sure
	default:
		return whatever
	}

}
