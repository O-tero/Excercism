package rotationalcipher

import "unicode"

func RotationalCipher(plain string, shiftKey int) string {
	result := make([]rune, len(plain))

	for i, char := range plain {
		// Check if the character is a letter
		if unicode.IsLetter(char) {
			// Determine the base ASCII value for lowercase or uppercase letters
			base := 'a'
			if unicode.IsUpper(char) {
				base = 'A'
			}
			// Apply the rotational shift using modular arithmetic
			shiftedChar := (char-base+rune(shiftKey))%26 + base
			result[i] = shiftedChar
		} else {
			// If the character is not a letter, keep it unchanged
			result[i] = char
		}
	}

	// Convert the slice of runes to a string and return it
	return string(result)
}
