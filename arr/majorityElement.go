package arr

// MajorityElement
// https://redirect.liyou-chen.com/mc8gmk

func MajorityElement(nums []int) int {
	counter := make(map[int]int)

	size := len(nums)
	for _, val := range nums {
		counter[val]++
		if counter[val] > size/2 {
			return val
		}
	}
	return 0
}
