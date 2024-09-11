package summultiples


func SumMultiples(limit int, divisors ...int) int {
	// Use a map to track unique multiples
	multiples := make(map[int]struct{})

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}

		for multiple := divisor; multiple < limit; multiple += divisor {
			multiples[multiple] = struct{}{}
		}
	}

	sum := 0
	for multiple := range multiples {
		sum += multiple
	}

	return sum
}
