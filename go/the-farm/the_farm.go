package thefarm

import "fmt"
import "errors"


type InvalidCowsError struct {
	cows    int
	message string
}

func (e *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", e.cows, e.message)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	fodder, err := fodderCalculator.FodderAmount(cows)
	if err != nil {
		return 0, err
	}
	fatteningFactor, err := fodderCalculator.FatteningFactor()
	if err != nil {
		return 0, err
	}
	return fodder * fatteningFactor / float64(cows), nil
}

func ValidateInputAndDivideFood(fodderCalculator FodderCalculator, cows int) (float64, error) {
	if cows <= 0 {
		return 0, errors.New("invalid number of cows")
	}
	return DivideFood(fodderCalculator, cows)
}

func ValidateNumberOfCows(cows int) error {
	switch {
	case cows < 0:
		return &InvalidCowsError{
			message: "there are no negative cows",
			cows:    cows,
		}
	case cows == 0:
		return &InvalidCowsError{
			message: "no cows don't need food",
			cows:    cows,
		}
	}
	return nil
}