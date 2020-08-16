package program

func maxProfit(prices []int) int {
	var res int
	for i := 0; i < len(prices)-1; i++ {
		if prices[i] < prices[i+1] {
			res += prices[i+1] - prices[i]
		}
	}
	return res
}
