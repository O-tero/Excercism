package piglatin

import (
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Fields(sentence) 
	var pigLatinWords []string
	for _, word := range words {
		pigLatinWords = append(pigLatinWords, processWord(word))
	}
	return strings.Join(pigLatinWords, " ") 
}

func processWord(word string) string {
	if startsWithVowel(word) {
		return word + "ay"
	}
	if startsWithQu(word) {
		return word[2:] + "quay"
	}
	if hasQu(word) {
		index := strings.Index(word, "qu")
		return word[index+2:] + word[:index] + "quay"
	}
	return handleConsonants(word)
}

func startsWithVowel(word string) bool {
	vowels := "aeiou"
	return strings.HasPrefix(word, "xr") ||
		strings.HasPrefix(word, "yt") ||
		strings.ContainsAny(string(word[0]), vowels)
}

func startsWithQu(word string) bool {
	return strings.HasPrefix(word, "qu")
}

func hasQu(word string) bool {
	return strings.Contains(word, "qu")
}

func handleConsonants(word string) string {
	vowels := "aeiou"
	for i, char := range word {
		// Treat 'y' as a vowel if it is not at the start of the word
		if strings.ContainsRune(vowels, char) || (char == 'y' && i > 0) {
			return word[i:] + word[:i] + "ay"
		}
	}
	return word + "ay" // Fallback
}


func startsWithConsonantFollowedByQu(word string) bool {
	return strings.Contains(word, "qu") && strings.Index(word, "qu") == len(getLeadingConsonants(word))
}

func findFirstVowel(word string) int {
	for i, char := range word {
		if strings.ContainsAny(string(char), "aeiou") {
			return i
		}
	}
	return -1
}

func getLeadingConsonants(word string) string {
	var consonants string
	for _, char := range word {
		if strings.ContainsAny(string(char), "aeiou") {
			break
		}
		consonants += string(char)
	}
	return consonants
}
