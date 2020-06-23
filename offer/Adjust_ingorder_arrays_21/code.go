package Adjust_ingorder_arrays_21

func exchange(nums []int) []int {
	nums_len := len(nums)
	if nums_len ==0 || nums_len ==1 {
		return nums
	}
	l, r := 0, nums_len-1

	for l<r {
		if nums[l]%2==1 && nums[r]%2==1 {
			l++
		} else if nums[l]%2==1 && nums[r]%2==0 {
			l++
			r--
		} else if nums[l]%2==0 && nums[r]%2==0 {
			r--
		} else if nums[l]%2==0 && nums[r]%2==1 {
			nums[l], nums[r] = nums[r], nums[l]
			l++
			r--
		}
	}
	return nums
}