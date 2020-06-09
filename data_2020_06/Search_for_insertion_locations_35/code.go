package Search_for_insertion_locations_35

func searchInsert(nums []int, target int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}
	for i := 0; i < l; i++ {
		if nums[i] == target || nums[i] > target {
			return i
		}
	}
	return l

}