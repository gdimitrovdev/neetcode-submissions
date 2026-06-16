func maxProfit(prices []int) int {
	var prevPrice int = prices[0]
	var profit int = 0

	for i := 0; i < len(prices); i++ {
		if prices[i] < prevPrice {
			prevPrice = prices[i]
		}

		if prices[i] > prevPrice {
			profit += prices[i] - prevPrice
			prevPrice = prices[i]
		}
	}

	return profit
}
