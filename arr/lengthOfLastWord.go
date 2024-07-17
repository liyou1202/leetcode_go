package arr

import (
	"regexp"
)

// LengthOfLastWord
// https://redirect.liyou-chen.com/q076y1
func LengthOfLastWord(s string) int {
	re := regexp.MustCompile("([a-zA-Z]+)")
	result := re.FindAllString(s, -1)
	return len(result[len(result)-1])
}
