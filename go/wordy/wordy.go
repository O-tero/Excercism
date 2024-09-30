package wordy

import (
	"regexp"
	"strconv"
)

func Answer(question string) (int, bool) {
	pattern := `^What is (-?\d+)((?: (plus|minus|multiplied by|divided by) (-?\d+))*)\?$`
	re := regexp.MustCompile(pattern)

	match := re.FindStringSubmatch(question)
	if match == nil {
		return 0, false // Invalid structure
	}

	result, err := strconv.Atoi(match[1])
	if err != nil {
		return 0, false
	}

	operations := match[2]
	if operations != "" {
		// Further match individual operations and numbers using another regex.
		opPattern := ` (plus|minus|multiplied by|divided by) (-?\d+)`
		opRe := regexp.MustCompile(opPattern)

		opMatches := opRe.FindAllStringSubmatch(operations, -1)
		for _, opMatch := range opMatches {
			operation := opMatch[1]
			number, err := strconv.Atoi(opMatch[2])
			if err != nil {
				return 0, false
			}

			switch operation {
			case "plus":
				result += number
			case "minus":
				result -= number
			case "multiplied by":
				result *= number
			case "divided by":
				if number == 0 {
					return 0, false
				}
				result /= number
			default:
				return 0, false 
			}
		}
	}

	return result, true
}
