package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	nums1 := []int{1, 1, 1, 2, 2, 3}
	k := arr.RemoveDuplicates2(nums1)
	fmt.Print(nums1, k)
}
