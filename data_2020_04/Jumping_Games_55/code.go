package Jumping_Games_55

func canJump(nums []int) bool {
	max_i := 0
	for i, v := range nums {
		if max_i >= i && i+v> max_i {
			max_i = i+v
		}
	}
	return max_i >= len(nums)-1
}