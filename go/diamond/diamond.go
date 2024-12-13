package diamond

import (
	"errors"
	"strings"
)

func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z' {
		return "", errors.New("input must be an uppercase letter between A and Z")
	}

	index := int(char - 'A') 
	var rows []string        

	for i := 0; i <= index; i++ {
		leadingSpaces := strings.Repeat(" ", index-i)
		letter := string(byte('A') + byte(i)) 
		if i == 0 {
			rows = append(rows, leadingSpaces+letter+leadingSpaces)
		} else {
			innerSpaces := strings.Repeat(" ", 2*i-1)
			rows = append(rows, leadingSpaces+letter+innerSpaces+letter+leadingSpaces)
		}
	}

	for i := index - 1; i >= 0; i-- {
		rows = append(rows, rows[i])
	}

	result := strings.Join(rows, "\n")
	return result, nil
}
