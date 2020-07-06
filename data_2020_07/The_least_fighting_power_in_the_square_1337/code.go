package The_least_fighting_power_in_the_square_1337

import (
	"sort"
)

func kWeakestRows(mat [][]int, k int) []int {
	res := []int{}
	for i, v := range mat {
		res[i] = sumsum(v)*100+i
	}
	sort.Ints(res)
	for i:=0;i<k;i++{
		res[i] = res[i]%100
	}
	return res[:k]
}

func sumsum(nums []int) int {
	res := 0
	for _, num := range nums {
		res += num
	}
	return res
}