// Package twelve returns a single verse for any of the twelve days of Christmas
package twelve

import (
	"fmt"
)

const testVersion = 1

func Verse(i int) string {
	fmt.Println("this is here")
	return "oh ya"
}

func Song() string {
	fmt.Println("this is that")
	return "string"
}
