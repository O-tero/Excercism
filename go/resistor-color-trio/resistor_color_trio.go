package resistorcolortrio

import "fmt"

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with a unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
    colorMap := map[string]int{
        "black": 0,
        "brown": 1,
        "red": 2,
        "orange": 3,
        "yellow": 4,
        "green": 5,
        "blue": 6,
        "violet": 7,
        "grey": 8,
        "white": 9,
    }

    baseValue := colorMap[colors[0]]*10 + colorMap[colors[1]]
    multiplier := colorMap[colors[2]]

    resistance := baseValue * intPow(10, multiplier)

    switch {
    case resistance >= 1_000_000_000:
        return fmt.Sprintf("%d gigaohms", resistance/1_000_000_000)
    case resistance >= 1_000_000:
        return fmt.Sprintf("%d megaohms", resistance/1_000_000)
    case resistance >= 1_000:
        return fmt.Sprintf("%d kiloohms", resistance/1_000)
    default:
        return fmt.Sprintf("%d ohms", resistance)
    }
}

func intPow(base, exp int) int {
    result := 1
    for i := 0; i < exp; i++ {
        result *= base
    }
    return result
}
