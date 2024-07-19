package arr

import (
	"regexp"
	"strings"
)

// IsPalindrome
// https://redirect.liyou-chen.com/o6ukue
func IsPalindrome(s string) bool {
	// 去除所有符號及空格
	re := regexp.MustCompile("([a-zA-Z0-9]+)")
	pureStrs := re.FindAllString(s, -1)
	var str = ""
	for _, val := range pureStrs {
		str += val
	}
	str = strings.ToLower(str)

	i, j := 0, len(str)-1

	for i < j {
		if str[i] != str[j] {
			return false
		}
		i++
		j--
	}

	return true
}
