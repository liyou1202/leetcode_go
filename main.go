package main

import (
	"fmt"
	"leetcodePractice/intervals"
)

func main() {
	input := []int{
		0, 1, 2, 4, 5, 7,
	}
	result := intervals.SummaryRanges(input)
	fmt.Println(result)
}
