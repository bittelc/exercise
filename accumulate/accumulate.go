// Package accumulate implements a given function provided to the package, against a series of arguments provided to the package
package accumulate

import ()

const testVersion = 1

func Accumulate(input []string, fn func(mod string) string) []string {
	var applied []string
	for _, str := range input {
		applied = append(applied, fn(str))
	}
	return applied
}
