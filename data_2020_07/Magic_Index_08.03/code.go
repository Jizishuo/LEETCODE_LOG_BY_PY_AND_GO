package Magic_Index_08_03

func findMagicIndex(nums []int) int {
	if len(nums)==0{
		return -1
	}
	if len(nums)==1 {
		if nums[0]==0{
			return 0
		} else {
			return -1
		}
	}

	l, r := 0, len(nums)
	for l< r {
		if nums[l]>l {
			l = nums[l]
		} else if nums[l]==l {
			return l
		} else {
			l++
		}
	}
	return -1
}