package Repeated_numbers_in_the_array_03

import (
	"sort"
)

func findRepeatNumber(nums []int) int {
	for i, v := range nums {
		if i == v {
			continue
		}
		if nums[i] == v {
			return v
		}
		nums[i], nums[v] = nums[v], nums[i]
	}
	return 0
}