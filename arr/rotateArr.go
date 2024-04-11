package arr

// Rotate
// https://redirect.liyou-chen.com/7kfu5y
func Rotate(nums []int, k int) {
	k = k % len(nums)
	reverse(nums)
	reverse(nums[:k])
	reverse(nums[k:])
}

func reverse(arr []int) {
	start, end := 0, len(arr)-1

	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}
