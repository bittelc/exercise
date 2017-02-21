// Package acronym converts a given string into its acronym
// Eg. "Is this a sentence" => "ITAS"
package acronym

import "strings"
import "fmt"

const testVersion = 2

func Abbreviate(str string) string {
	var acronym string
	words := strings.Fields(str)
	for _, word := range words {
		acronym += fmt.Sprintf("%c", word[0])
	}
	return acronym
}
