package encode

import (
	"strconv"
	"strings"
)


func RunLengthEncode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	count := 1
	current := input[0]

	for i := 1; i < len(input); i++ {
		if input[i] == current {
			count++
		} else {
			if count > 1 {
				result.WriteString(strconv.Itoa(count))
			}
			result.WriteByte(current) 
			current = input[i]       
			count = 1               
		}
	}

	
	if count > 1 {
		result.WriteString(strconv.Itoa(count))
	}
	result.WriteByte(current)

	return result.String()
}


func RunLengthDecode(input string) string {
	if len(input) == 0 {
		return ""
	}

	var result strings.Builder
	var countStr strings.Builder

	for _, char := range input {
		if char >= '0' && char <= '9' { 
			countStr.WriteRune(char)
		} else {
			
			count := 1
			if countStr.Len() > 0 {
				count, _ = strconv.Atoi(countStr.String()) 
				countStr.Reset()
			}

			for i := 0; i < count; i++ { 
				result.WriteRune(char)
			}
		}
	}

	return result.String()
}
