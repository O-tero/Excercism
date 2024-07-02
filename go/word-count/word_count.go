package wordcount

import (
    "regexp"
    "strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
    phrase = strings.ToLower(phrase)

    re := regexp.MustCompile(`\b\w+('\w+)?\b`)
    words := re.FindAllString(phrase, -1)
    freq := make(Frequency)
    for _, word := range words {
        freq[word]++
    }

    return freq
}
