package lsproduct

import (
    "errors"
    "strconv"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
    if span > len(digits) || span < 0 {
        return 0, errors.New("invalid span")
    }
    // Edge case: span is zero, product is 1
    if span == 0 {
        return 1, nil
    }

    numDigits := make([]int, len(digits))
    for i, char := range digits {
        digit, err := strconv.Atoi(string(char))
        if err != nil {
            return 0, errors.New("invalid input")
        }
        numDigits[i] = digit
    }

    var maxProduct int64 = 0
    for i := 0; i <= len(numDigits)-span; i++ {
        var product int64 = 1
        for j := i; j < i+span; j++ {
            product *= int64(numDigits[j])
        }
        if product > maxProduct {
            maxProduct = product
        }
    }

    return maxProduct, nil
}