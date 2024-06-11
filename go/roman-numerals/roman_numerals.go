package romannumerals

import "errors"

func ToRomanNumeral(input int) (string, error) {
		if input <= 0 || input > 3999 {
		return "", errors.New("invalid number, please enter a number between 1 and 3999")
	}

	// Mapping of Arabic numeral values to Roman numerals
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	numerals := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	// Construct the Roman numeral string
	roman := ""
	for i := 0; i < len(values); i++ {
		for input >= values[i] {
			input -= values[i]
			roman += numerals[i]
		}
	}

	return roman, nil
}
