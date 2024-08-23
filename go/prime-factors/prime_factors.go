package prime

func Factors(n int64) []int64 {
	var factors []int64
    var divisor int64 = 2

    for n > 1 {
        for n%divisor == 0 {
            factors = append(factors, divisor)
            n /= divisor
        }
        divisor++
    }

    return factors
}
