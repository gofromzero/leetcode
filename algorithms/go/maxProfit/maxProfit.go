package maxProfit

import "math"

func maxProfit(prices []int) int {
	var maxprofit int
	var minPrice = math.MaxInt
	for i := 0; i < len(prices); i++ {
		if prices[i] < minPrice {
			minPrice = prices[i]
		} else if prices[i]-minPrice > maxprofit {
			maxprofit = prices[i] - minPrice
		}
	}

	return maxprofit
}
