package summultiples

// SumMultiples calculates the sum of unique multiples of the divisors below the limit.
func SumMultiples(limit int, divisors ...int) int {
	sum := 0
	visited := make([]bool, limit)

	for _, divisor := range divisors {
		if divisor == 0 {
			continue
		}

		for multiple := divisor; multiple < limit; multiple += divisor {
			if !visited[multiple] {
				visited[multiple] = true
				sum += multiple
			}
		}
	}

	return sum
}
