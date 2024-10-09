package cryptosquare

import (
	"math"
	"strings"
)

func Encode(s string) string {
	var runes, out []rune
	for _, r := range strings.ToLower(s) {
		if !(r < '0' || (r > '9' && r < 'a') || r > 'z') {
			runes = append(runes, r)
		}
	}
	len := len(runes)
	cols := int(math.Ceil(math.Sqrt(float64(len))))
	for i := 0; i < cols; i++ {
		if i != 0 {
			out = append(out, ' ')
		}
		for j := 0; j < len; j += cols {
			if i+j < len {
				out = append(out, runes[i+j])
			} else {
				out = append(out, ' ')
			}
		}
	}
	return string(out)
}