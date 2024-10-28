package allyourbase

import (
	"errors"
	"math"
)


func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}


	if len(inputDigits) == 0 || isAllZeros(inputDigits) {
		return []int{0}, nil
	}

	
	numberInBase10 := baseToDecimal(inputBase, inputDigits)


	outputDigits := decimalToBase(outputBase, numberInBase10)

	return outputDigits, nil
}

func isAllZeros(digits []int) bool {
	for _, digit := range digits {
		if digit != 0 {
			return false
		}
	}
	return true
}


func baseToDecimal(base int, digits []int) int {
	number := 0
	for i, digit := range digits {
		power := len(digits) - 1 - i
		number += digit * int(math.Pow(float64(base), float64(power)))
	}
	return number
}


func decimalToBase(base int, number int) []int {
	if number == 0 {
		return []int{0}
	}

	var digits []int
	for number > 0 {
		remainder := number % base
		digits = append([]int{remainder}, digits...)
		number /= base
	}

	return digits
}
