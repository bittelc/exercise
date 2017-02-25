// Package acronym converts a given string into its acronym
// Eg. "Is this a sentence" => "ITAS"
package acronym

import (
	"fmt"
	"regexp"
	"strings"
)

const testVersion = 2

func Abbreviate(str string) string {
	criteria := regexp.MustCompile("[A-Z]")
	regex_match := criteria.FindAllString(str, -1)
	fmt.Printf(strings.Join(regex_match[:], ""))
	// return fmt.Sprintf(strings.TrimSpace(fmt.Sprintf("%v", regex_match)))
	return strings.Join(regex_match[:], "")
}
