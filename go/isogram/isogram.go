package isogram

import "unicode"

// IsIsogram checks if a word or phrase is an isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, char := range word {
		if unicode.IsLetter(char) {
			lowerChar := unicode.ToLower(char)
			if seen[lowerChar] {
				return false
			}
			seen[lowerChar] = true
		}
	}
	return true
}
