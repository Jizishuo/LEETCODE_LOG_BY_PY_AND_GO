package The_total_assets_of_the_richest_clients_1672

func maximumWealth(accounts [][]int) int {
	var max_u int
	for i:=0;i<len(accounts);i++ {
		var b int
		b = sum(accounts[i])
		if b > max_u {
			max_u = b
		}
	}
	return max_u
}

func sum(l []int) int {
	a := 0
	for _, v := range l {
		a += v
	}
	return a
}
