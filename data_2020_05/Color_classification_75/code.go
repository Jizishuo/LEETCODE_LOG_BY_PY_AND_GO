package Color_classification_75

func sortColors(nums []int)  {
	l := 0
	r := len(nums)-1
	for i:=0;i<=r;i++ {
		if nums[i]==0{
			nums[i], nums[l]=nums[l], nums[i]
			l++
		}
		if nums[i]==2{
			nums[i], nums[r]=nums[r], nums[i]
			i--
			r--
		}
	}


}