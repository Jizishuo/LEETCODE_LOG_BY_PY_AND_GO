package Removes_duplicates_from_the_sorted_array_26

func removeDuplicates(nums []int) int {
	for i := 0; i+1 < len(nums); {
		if nums[i] == nums[i+1] {
			nums = append(num[:i], nums[i+1:]...)
		} else {
			i++
		}
	}
	return len(nums)

}

