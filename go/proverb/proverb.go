package proverb

import "fmt"

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
    var result []string

	if len(rhyme) == 0 {
		return result
	}
	for i := 0; i < len(rhyme)-1; i++ {
		line := fmt.Sprintf("For want of a %s the %s was lost.", rhyme[i], rhyme[i+1])
		result = append(result, line)
	}
	finalLine := fmt.Sprintf("And all for the want of a %s.", rhyme[0])
	result = append(result, finalLine)

	return result
}
