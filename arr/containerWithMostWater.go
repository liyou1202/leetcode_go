package arr

// MaxArea
// https://redirect.liyou-chen.com/fwsfaz
func MaxArea(height []int) int {
	var i, j = 0, len(height) - 1

	maximum := 0
	for i < j {
		y := j - i
		x := 0
		if height[i] >= height[j] {
			x = height[j]
			j--
		} else {
			x = height[i]
			i++
		}

		temp := x * y
		if temp > maximum {
			maximum = temp
		}
	}

	return maximum
}
