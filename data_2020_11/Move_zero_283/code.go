package Move_zero_283

func moveZeroes(nums []int)  {

	i_0 := 0

	for i:=0;i<len(nums);i++ {
		if nums[i] == 0 {
			continue
		}
		if i_0 == i {
			i_0 ++
			continue
		}
		nums[i_0] = nums[i]
		nums[i] = 0
		i_0 ++
	}
}