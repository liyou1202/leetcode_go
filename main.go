package main

import (
	"fmt"
	"leetcodePractice/arr"
)

func main() {
	heightArr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	area := arr.MaxArea(heightArr)
	fmt.Println(area)
}
