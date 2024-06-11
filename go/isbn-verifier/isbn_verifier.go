package isbn

import (
	"strings"
    "unicode"
)

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
	isbn = strings.ReplaceAll(isbn, " ", "")
	
	// Check the length of the ISBN-10 string
	if len(isbn) != 10 {
		return false
	}

	total := 0
	for i, char := range isbn {
		// Special case for 'X' which represents 10
		if char == 'X' {
			if i != 9 {
				return false
			}
			total += 10
		} else if unicode.IsDigit(char) {
			total += int(char-'0') * (10 - i)
		} else {
			return false
		}
	}

	return total%11 == 0
}
