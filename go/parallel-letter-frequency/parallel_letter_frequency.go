package letter


type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(text string) FreqMap {
    frequencies := FreqMap{}
    for _, r := range text {
        frequencies[r]++
    }
    return frequencies
}

// ConcurrentFrequency counts the frequency of each rune in the given texts
// by making use of concurrency.
func ConcurrentFrequency(texts []string) FreqMap {
    ch := make(chan FreqMap)
    
    for _, text := range texts {
        go func(t string) {
            ch <- Frequency(t)
        }(text)
    }
    
    result := FreqMap{}
    
    for range texts {
        freqMap := <-ch
        for char, count := range freqMap {
            result[char] += count
        }
    }
    
    return result
}
