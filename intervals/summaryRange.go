package intervals

import (
	"fmt"
	"strconv"
)

// SummaryRanges
// https://redirect.liyou-chen.com/4mvfvs
func SummaryRanges(nums []int) []string {
	var result []string
	isContinue := false
	var start int
	for i := 0; i < len(nums); i++ {
		if i == len(nums)-1 || nums[i]+1 != nums[i+1] {
			if !isContinue {
				result = append(result, strconv.Itoa(nums[i]))
				continue
			}
			//組合String
			combineStr := fmt.Sprintf("%d->%d", start, nums[i])
			result = append(result, combineStr)
			isContinue = false
			continue
		}
		if !isContinue {
			start = nums[i]
		}
		isContinue = true
	}
	return result
}
