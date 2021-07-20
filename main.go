package main

func maxProfit(prices []int) int {
	min := 999999999
	var max int

	for i := 0; i < len(prices); i++ {
		cur := prices[i]
		if cur <= min {
			min = cur
		} else if (cur - min) > max {
			max = cur - min
		}
	}
	return max
}
