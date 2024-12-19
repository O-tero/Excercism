package yacht

import (
	"sort"
	"strings"
)

func Score(dice []int, category string) int {
	category = strings.ToLower(category)

	counts := make(map[int]int)
	for _, die := range dice {
		counts[die]++
	}

	switch category {
	case "ones", "twos", "threes", "fours", "fives", "sixes":
		target := map[string]int{
			"ones": 1, "twos": 2, "threes": 3,
			"fours": 4, "fives": 5, "sixes": 6,
		}[category]
		return target * counts[target]

	case "full house":
		three, two := false, false
		total := 0
		for num, count := range counts {
			if count == 3 {
				three = true
			} else if count == 2 {
				two = true
			}
			total += num * count
		}
		if three && two {
			return total
		}
		return 0

	case "four of a kind":
		for num, count := range counts {
			if count >= 4 {
				return num * 4
			}
		}
		return 0

	case "little straight":
		expected := []int{1, 2, 3, 4, 5}
		sort.Ints(dice)
		if equal(dice, expected) {
			return 30
		}
		return 0

	case "big straight":
		expected := []int{2, 3, 4, 5, 6}
		sort.Ints(dice)
		if equal(dice, expected) {
			return 30
		}
		return 0

	case "choice":	
		total := 0
		for _, die := range dice {
			total += die
		}
		return total

	case "yacht":
		for _, count := range counts {
			if count == 5 {
				return 50
			}
		}
		return 0
	}

	return 0
}

// Helper function to check equality of two sorted slices
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
