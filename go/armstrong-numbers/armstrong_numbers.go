package armstrong

import "math"

func IsNumber(n int) bool {
	originalNumber := n
	numDigits := 0

	for temp := n; temp > 0; temp /= 10 {
		numDigits++
	}

	sum := 0
    
	for temp := n; temp > 0; temp /= 10 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(numDigits)))
	}

	return sum == originalNumber
}