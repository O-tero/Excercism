package triangle

import (
	"math"
)

// Kind represents the type of a triangle.
type Kind int

const (
	NaT Kind = iota // Not a triangle
	Equ             // Equilateral
	Iso             // Isosceles
	Sca             // Scalene
)

// KindFromSides determines the type of a triangle based on its side lengths.
func KindFromSides(a, b, c float64) Kind {
	if a <= 0 || b <= 0 || c <= 0 || math.IsNaN(a) || math.IsNaN(b) || math.IsNaN(c) || math.IsInf(a, 0) || math.IsInf(b, 0) || math.IsInf(c, 0) {
		return NaT
	}

	// Check the triangle inequality
	if a+b < c || a+c < b || b+c < a {
		return NaT
	}

	// Determine the type of triangle
	if a == b && b == c {
		return Equ
	} else if a == b || b == c || a == c {
		return Iso
	} else {
		return Sca
	}
}
