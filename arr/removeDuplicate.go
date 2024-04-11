package arr

// RemoveDuplicates
// https://redirect.liyou-chen.com/tt311s
func RemoveDuplicates(nums []int) int {
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
		}
	}
	return count
}
