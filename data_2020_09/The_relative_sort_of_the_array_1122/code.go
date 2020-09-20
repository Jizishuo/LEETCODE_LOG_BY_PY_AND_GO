package The_relative_sort_of_the_array_1122

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	lastidx :=0
	for i:=0;i<len(arr2);i++ {
		for j:= lastidx;j<len(arr1);j++ {
			if arr1[j]==arr2[i] {
				arr1[lastidx], arr1[j] = arr1[j], arr1[lastidx]
				lastidx ++
			}
		}

	}
	sort.Ints(arr1[lastidx:])
	return arr1
}