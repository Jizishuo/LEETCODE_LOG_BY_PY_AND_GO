package Find_the_first_and_last_position_of_an_element_in_a_sort_array_34


func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	start, end := -1, -1
	for left<=right {
		mid := (left+right)/2
		if nums[mid]<target{
			left = mid +1
		} else if nums[mid] >target {
			right = mid -1
		} else {
			ll, lr := left, mid-1
			for ll <= lr {
				lm := (ll+lr)/2
				if nums[lm] == target {
					lr = lm -1
				} else {
					ll = lm +1
				}
			}
			rl, rr := mid +1, right
			for rl<=rr {
				rm := (rl+rr)/2
				if nums[rm] == target {
					rl = rm+1
				} else {
					rr = rm-1
				}
			}
			start, end = ll, rr
			break
		}
	}
	return []int{start, end}
}