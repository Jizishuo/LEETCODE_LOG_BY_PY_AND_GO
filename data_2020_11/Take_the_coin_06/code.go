package Take_the_coin_06

func minCount(coins []int) int {
	var res int
	for _, v := range coins {
		res += (v+1)/2
	}
	return res
}