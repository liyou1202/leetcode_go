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

// RemoveDuplicates2
// https://redirect.liyou-chen.com/fzbpny
func RemoveDuplicates2(nums []int) int {
	size := len(nums)
	if size <= 2 {
		return size
	}

	count := 1
	duplicateCount := 0
	for i := 1; i < size; i++ {
		if nums[i] != nums[i-1] {
			nums[count] = nums[i]
			count++
			duplicateCount = 0
			continue
		}

		if duplicateCount < 1 {
			nums[count] = nums[i]
			count++
			duplicateCount++
		}
	}
	return count
}
