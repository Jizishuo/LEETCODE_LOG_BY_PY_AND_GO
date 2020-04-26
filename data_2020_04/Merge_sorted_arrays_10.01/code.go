package Merge_sorted_arrays_10_01

import "sort"

func merge(A []int, m int, B []int, n int)  {
	copy(A[m:], B[:])
	sort.Ints(A)
}