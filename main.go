package main

func main() {
	nums1 := []int{7, 1, 5, 3, 6, 4}
	println(maxProfit(nums1))
}

func maxProfit(prices []int) int {
	profit := 0
	buy := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < buy {
			buy = prices[i]
			continue
		}
		temp := prices[i] - buy
		if profit > 0 {
			profit += temp
			buy = prices[i]
		}
	}
	return profit
}
