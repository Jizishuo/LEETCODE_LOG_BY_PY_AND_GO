package The_best_time_to_buy_and_sell_stocks_122

func maxProfit(prices []int) int {
	res := 0
	for i:=0;i<len(prices)-1;i++ {
		if prices[i+1]>prices[i] {
			res = res + prices[i+1]-prices[i]
		}
	}
	return res
}