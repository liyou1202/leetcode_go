package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	nums1 := []int{0, 1, 2, 2, 3, 0, 4, 2}

	k := arr.RemoveElement(nums1, 2)
	fmt.Print(nums1, k)
}
