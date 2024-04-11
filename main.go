package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	nums2 := []int{2, 5, 6}
	var m, n int = 3, 3
	arr.MergeSortedArr(nums1, m, nums2, n)
	fmt.Print(nums1)
}
