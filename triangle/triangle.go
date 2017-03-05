// Package triangle allows for determining whether a series of connected line lengths creates a triangle; and if so, what kind
package triangle

import (
	"math"
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

	lengths := []float64{a, b, c}
	sort.Float64s(lengths)
	short := lengths[0]
	mid := lengths[1]
	long := lengths[2]

	if short == 0 || mid == 0 || long == 0 ||
		short+mid < long ||
		math.IsNaN(short) ||
		math.IsInf(long, 0) {
		return NaT
	}

	if short == mid && mid == long {
		return Equ
	}

	if short == mid || mid == long {
		return Iso
	}

	return Sca
}
