package main

func maxProfit(prices []int) int {
	maxProfit := 0
	stack := []int{}

	for _, p := range prices {
		if p < stack[len(stack)-1] || len(stack) == 0 {
			stack = append(stack, p)
		} else {
			top := stack[len(stack)-1]
			profit := p - top
			if profit > maxProfit {
				maxProfit = profit
			}
		}
	}
	return maxProfit
}

func maxProfit2(prices []int) int {
	maxProfit := 0
	minPrice := prices[0]

	for i := 1; i < len(prices); i++ {
		minPrice = min(prices[i], minPrice)
		maxProfit = max(prices[i]-minPrice, maxProfit)
	}
	return maxProfit
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
