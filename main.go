package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	nums1 := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	arr.Rotate(nums1, k)
	fmt.Print(nums1)
}
