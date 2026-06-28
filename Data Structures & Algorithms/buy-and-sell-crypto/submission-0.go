func maxProfit(prices []int) int {
	minLeft := prices[0]
	maxProfit := 0

	for i := 0; i < len(prices); i++ {
		minLeft = min(minLeft, prices[i])
		prices[i] = prices[i] - minLeft
		maxProfit = max(maxProfit, prices[i])
	}

	return maxProfit
}
