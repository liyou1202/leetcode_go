package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	nums1 := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	k := arr.RemoveDuplicates(nums1)
	fmt.Print(nums1, k)
}
