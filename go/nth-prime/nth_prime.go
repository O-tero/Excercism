package prime

import (
    "errors"
    "math"
)

// isPrime checks if a number is prime.
func isPrime(num int) bool {
    if num <= 1 {
        return false
    }
    for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
        if num%i == 0 {
            return false
        }
    }
    return true
}

// NthPrime returns the nth prime number.
func Nth(n int) (int, error) {
    if n < 1 {
        return 0, errors.New("input must be greater than 0")
    }
    
    count := 0
    num := 1
    for count < n {
        num++
        if isPrime(num) {
            count++
        }
    }
    
    return num, nil
}
