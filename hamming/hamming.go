//Package hamming calculates the hamming distance between two homologuos strings.
package hamming

import "fmt"

const testVersion = 5

// Distance takes two strings and compares their constituent characters.
func Distance(a, b string) (int, error) {

	if len(a) != len(b) {
		return -1, fmt.Errorf("Unequal string lengths!")
	}

	mutations := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			mutations += 1
		}
	}

	return mutations, nil
}
