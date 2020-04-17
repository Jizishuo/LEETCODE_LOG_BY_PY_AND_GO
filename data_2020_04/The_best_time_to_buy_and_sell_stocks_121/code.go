package The_best_time_to_buy_and_sell_stocks_121

func maxProfit(prices []int) int {
	if len(prices) <=0 {
		return 0
	}
	minnum:=int(^uint(0)>>1)                  //go语言中的最大int值
	maxprice := 0
	for _, v := range prices {
		maxprice = max(v-minnum, maxprice)
		minnum = mix(v, minnum)
	}
	return maxprice
}

func mix(x,y int) int {
	if x>y {
		return y
	} else {
		return x
	}
}
func max(x,y int) int {
	if x>y {
		return x
	} else {
		return y
	}
}