package grep

import (
	"bufio"
	"os"
	"strings"
    "strconv"
)

func Search(pattern string, flags, files []string) []string {
	var result []string
	caseInsensitive := contains(flags, "-i")
	invertMatch := contains(flags, "-v")
	matchEntireLine := contains(flags, "-x")
	prependLineNumber := contains(flags, "-n")
	onlyFilenames := contains(flags, "-l")

	// Adjust pattern for case-insensitive search
	if caseInsensitive {
		pattern = strings.ToLower(pattern)
	}

	for _, file := range files {
		matches := []string{}
		fileHandle, err := os.Open(file)
		if err != nil {
			continue 
		}
		defer fileHandle.Close()

		scanner := bufio.NewScanner(fileHandle)
		lineNumber := 0

		for scanner.Scan() {
			lineNumber++
			line := scanner.Text()
			originalLine := line

			if caseInsensitive {
				line = strings.ToLower(line)
			}
            
			matchesPattern := strings.Contains(line, pattern)
			if matchEntireLine {
				matchesPattern = line == pattern
			}

			// Invert match if necessary
			if invertMatch {
				matchesPattern = !matchesPattern
			}

			if matchesPattern {
				if onlyFilenames {
					result = append(result, file)
					break
				}

				match := ""
				if prependLineNumber {
					match = formatMatchWithLineNumber(file, lineNumber, originalLine, len(files) > 1)
				} else if len(files) > 1 {
					match = file + ":" + originalLine
				} else {
					match = originalLine
				}
				matches = append(matches, match)
			}
		}

		if !onlyFilenames {
			result = append(result, matches...)
		}
	}

	return result
}

// Helper function to check if a slice contains a string
func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}

// Helper function to format a match with line number
func formatMatchWithLineNumber(file string, lineNumber int, line string, includeFilename bool) string {
	if includeFilename {
		return file + ":" + formatLineNumber(lineNumber) + line
	}
	return formatLineNumber(lineNumber) + line
}

// Helper function to prepend the line number
func formatLineNumber(lineNumber int) string {
	return strconv.Itoa(lineNumber) + ":"
}
