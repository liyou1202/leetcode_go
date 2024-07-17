package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	strs := []string{
		"a", "ab",
	}
	result := arr.LongestCommonPrefix(strs)
	fmt.Println(result)
}
