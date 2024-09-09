package bottlesong
import (
	"fmt"
	"strings"
)
var nums = []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}
func bottle(num int) string {
	if num == 1 {
		return "bottle"
	} else {
		return "bottles"
	}
}
func Recite(startBottles, takeDown int) []string {
	var result []string
	for i := takeDown; i > 0; i-- {
		s := fmt.Sprintf("%s green %s hanging on the wall,", strings.ToUpper(nums[startBottles][:1])+nums[startBottles][1:], bottle(startBottles))
		result = append(result, s)
		result = append(result, s)
		result = append(result, "And if one green bottle should accidentally fall,")
		result = append(result, fmt.Sprintf("There'll be %s green %s hanging on the wall.", nums[startBottles-1], bottle(startBottles-1)))
		result = append(result, "")
		startBottles--
	}
	return result[:len(result)-1]
}