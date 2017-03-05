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
	var triangle Kind
	sorted := sort.Float64s([]float64{a, b, c})
	fmt.Println(sorted)
	return triangle
}
