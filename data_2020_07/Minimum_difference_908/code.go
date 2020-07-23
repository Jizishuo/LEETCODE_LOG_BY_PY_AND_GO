package Minimum_difference_908

import (
	"math"
)

func smallestRangeI(A []int, K int) int {
	min,max := math.MaxInt64,math.MinInt64
	for i:=0;i<len(A);i++ {
		if A[i]>max {
			max=A[i]
		}
		if A[i]<min {
			min = A[i]
		}
	}
	res := (max-min)-2*K
	if  res<0 {
		return 0
	} else {
		return res
	}
}