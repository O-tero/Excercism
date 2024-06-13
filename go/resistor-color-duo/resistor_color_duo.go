package resistorcolorduo

import "strings"

// Color map based on resistor color codes
var colorMap = map[string]int{
	"black":  0,
	"brown":  1,
	"red":    2,
	"orange": 3,
	"yellow": 4,
	"green":  5,
	"blue":   6,
	"violet": 7,
	"grey":   8,
	"white":  9,
}

// Function to get the resistor value
func Value(colors []string) int {
	
	if len(colors) < 2 {
		return -1 
	}

	// Get the numeric values of the first two colors
	firstValue, exists1 := colorMap[strings.ToLower(colors[0])]
	secondValue, exists2 := colorMap[strings.ToLower(colors[1])]

	if !exists1 || !exists2 {
		return -1 
	}

	// Form the two-digit number
	result := firstValue*10 + secondValue

	return result
}
