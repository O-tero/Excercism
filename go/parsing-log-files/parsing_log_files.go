package parsinglogfiles

import (
	"regexp"
    "fmt"
)

func IsValidLine(text string) bool {
	pattern := `^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`
    regex := regexp.MustCompile(pattern)
    return regex.MatchString(text)
}

func SplitLogLine(text string) []string {
	r := regexp.MustCompile(`<[*~=-]*>`)
    return r.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	count := 0
    r := regexp.MustCompile(`(?i)".*password.*"`)
    for _, lines := range lines {
        if r.MatchString(lines) {
            count++
        }
    }
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line\d+`)
    return re.ReplaceAllString(text, "")
}

func TagWithUserName(lines []string) []string {
	var result []string
	r := regexp.MustCompile(`\s+User\s+(\S+)`)
	for _, line := range lines {
		match := r.FindStringSubmatch(line)
		if len(match) > 1 {
			line = fmt.Sprintf("[USR] %s %s", match[1], line)
		}
		result = append(result, line)
	}
	return result
}
