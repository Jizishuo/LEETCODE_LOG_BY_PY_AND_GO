package Find_the_numbers_that_disappear_from_all_arrays_448


func findDisappearedNumbers(nums []int) []int {
	result := make([]int, 0,0)
	for _, value := range nums {
		if value <0 {
			value = -value
		}
		if nums[value-1]>0 {
			nums[value-1] = -nums[value-1]
		}
	}
	for i, v := range nums {
		if v>0 {
			result =append(result, i+1)
		}
	}
	return result
}