package intervals

import (
	"sort"
)

// Merge
// https://leetcode.com/problems/merge-intervals/submissions/1326200701/?envType=study-plan-v2&envId=top-interview-150
func Merge(intervals [][]int) [][]int {
	if len(intervals) == 0 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var result [][]int
	current := intervals[0]
	for _, val := range intervals[1:] {
		if current[1] >= val[0] {
			if val[1] > current[1] {
				current[1] = val[1]
			}
			continue
		}

		result = append(result, current)
		current = val
	}
	result = append(result, current)

	return result
}
