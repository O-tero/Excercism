package perfect

import "errors"

// Define the Classification type
type Classification int

const (
	ClassificationDeficient Classification = iota
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("only positive integers are allowed")

func Classify(n int64) (Classification, error) {
	if n <= 0 {
		return 0, ErrOnlyPositive
	}

	aliquotSum := int64(0)
	for i := int64(1); i <= n/2; i++ {
		if n%i == 0 {
			aliquotSum += i
		}
	}

	switch {
	case aliquotSum == n:
		return ClassificationPerfect, nil
	case aliquotSum > n:
		return ClassificationAbundant, nil
	default:
		return ClassificationDeficient, nil
	}
}
