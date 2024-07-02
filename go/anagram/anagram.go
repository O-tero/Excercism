package anagram

import (
	"sort"
	"strings"
)

func normalize(s string) string {
	s = strings.ToLower(s)
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func Detect(subject string, candidates []string) []string {
	normalizedSubject := normalize(subject)
	var result []string

	for _, candidate := range candidates {
		if strings.ToLower(candidate) == strings.ToLower(subject) {
			continue
		}
		if normalize(candidate) == normalizedSubject {
			result = append(result, candidate)
		}
	}

	return result
}