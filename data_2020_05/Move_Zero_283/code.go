package Move_Zero_283

func moveZeroes(nums []int)  {
	count := 0
	for i:=0; i<len(nums); i ++ {
		if nums[i] != 0 {
			nums[i], nums[count] = nums[count], nums[i]
			count ++
		}
	}
}
