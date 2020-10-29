package Rearrange_the_arrays_1470

func shuffle(nums []int, n int) []int {
	res := make([]int, 0, n*2)
	for i:=0;i<n;i++ {
		res=append(res, nums[i], nums[i+n])
	}
	return res
}