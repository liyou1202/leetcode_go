package arr

// MergeSortedArr
// https://redirect.liyou-chen.com/vif85h
func MergeSortedArr(nums1 []int, m int, nums2 []int, n int) {
	i := m - 1
	j := n - 1

	if n == 0 {
		return
	}

	for k := m + n - 1; k >= 0; k-- {
		if j < 0 {
			return
		}

		if i < 0 {
			nums1[k] = nums2[j]
			j--
			continue
		}

		if nums1[i] >= nums2[j] {
			nums1[k] = nums1[i]
			i--
			continue
		}
		nums1[k] = nums2[j]
		j--
	}
}
