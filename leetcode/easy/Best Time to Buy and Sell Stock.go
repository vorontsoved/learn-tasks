package easy

func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxPr := 0

	for _, v := range prices {
		if v < minPrice {
			minPrice = v
		}
		cur := v - minPrice
		if maxPr < cur {
			maxPr = cur
		}
	}

	return maxPr
}
