// Package triangle allows for determining whether a series of connected line lengths creates a triangle; and if so, what kind
package triangle

import (
	"fmt"
	"sort"
)

const testVersion = 3

type Kind struct {
	NaT, Equ, Iso, Sca bool
}

// Pick values for the following identifiers used by the test program.
var NaT = Kind{NaT: true}
var Equ = Kind{Equ: true}
var Iso = Kind{Iso: true}
var Sca = Kind{Sca: true}

// KindFromSides accepts lenghts of sides of a triangle and classifies the triangle
func KindFromSides(a, b, c float64) Kind {
	triangle := NaT

	lengths := []float64{a, b, c}
	sort.Float64s(lengths)
	short := lengths[0]
	mid := lengths[1]
	long := lengths[2]

	if short == mid && mid == long {
		triangle = Equ
	}

	fmt.Println(lengths)
	fmt.Println(short)
	fmt.Println(mid)
	fmt.Println(long)

	return triangle
}
