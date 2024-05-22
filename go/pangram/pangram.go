package pangram

import (
    "strings"
    "unicode"
)

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	letterSet := make(map[rune]struct{})

    for _, char := range input{
        if unicode.IsLetter(char){
            letterSet[char] = struct{}{}
        }
    }
    return len(letterSet) == 26
}
