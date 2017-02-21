// Package acronym converts a given string into its acronym
// Eg. "Is this a sentence" => "ITAS"
package acronym

import "strings"
import "fmt"
import "unicode"

const testVersion = 2

func Abbreviate(str string) string {
	var acronym string
	f := func(c rune) bool {
		return !unicode.IsLetter(c)
	}
	words := strings.FieldsFunc(str, f)
	for _, word := range words {
		acronym += fmt.Sprintf("%c", word[0])
	}
	return acronym
}
