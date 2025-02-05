package pythagorean

import "math"

type Triplet [3]int

func Range(min, max int) []Triplet {
	var triplets []Triplet

	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c := int(math.Sqrt(float64(a*a + b*b)))
			if c > max { 
				break
			}
			if a*a+b*b == c*c { 
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}

// Sum finds all Pythagorean triplets where a + b + c equals the given sum.
func Sum(p int) []Triplet {
	var triplets []Triplet

	for a := 1; a < p/2; a++ { // a should be less than half of p
		for b := a + 1; b < p/2; b++ {
			c := p - a - b // Since a + b + c = p, we can directly compute c
			if a*a+b*b == c*c && a < b && b < c {
				triplets = append(triplets, Triplet{a, b, c})
			}
		}
	}
	return triplets
}
