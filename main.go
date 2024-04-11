package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	nums1 := []int{2, 2, 1, 1, 1, 2, 2}
	k := arr.MajorityElement(nums1)
	fmt.Print(nums1, k)
}
