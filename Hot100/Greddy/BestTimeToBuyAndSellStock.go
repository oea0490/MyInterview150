package Greddy

func maxProfit(prices []int) int {
	preMin := prices[0]
	maxProfit := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] < preMin {
			preMin = prices[i]
		} else {
			maxProfit = max(maxProfit, prices[i]-preMin)
		}
	}
	return maxProfit
}
