package Create_an_array_of_targets_in_a_given_order_1386

func createTargetArray(nums []int, index []int) []int {
	var res = make([]int, len(nums))
	for k,i := range index {
		copy(res[i+1:], res[i:])
		res[i] = nums[k]
	}
	return res
}