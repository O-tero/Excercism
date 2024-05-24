package reverse

func Reverse(input string) string {
	runes := []rune(input)
	n := len(runes)
	// Reverse the rune slice
	for i := 0; i < n/2; i++ {
		runes[i], runes[n-1-i] = runes[n-1-i], runes[i]
	}
	return string(runes)
}