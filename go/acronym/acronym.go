package acronym

import (
	"strings"
	"unicode"
)

func Abbreviate(s string) string {
	// Split the string using spaces or hyphens as delimiters
	words := strings.FieldsFunc(s, func(r rune) bool {
		return r == ' ' || r == '-'
	})
	// Collect the first letters of each word, ignoring punctuation
	var acronym string
	for _, word := range words {
	// Find the first letter in the word that is a letter (ignore punctuation)
		for _, char := range word {
			if unicode.IsLetter(char) {
				acronym += string(unicode.ToUpper(char))
				break
			}
		}
	}

	return acronym
}
