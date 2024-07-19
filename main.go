package main

import (
	"fmt"
	"leetcodePractice/intervals"
)

func main() {
	input := [][]int{
		{1, 4},
		{0, 4},
	}

	result := intervals.Merge(input)
	fmt.Println(result)
}
