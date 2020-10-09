package Dynamics_and_of_one_dimensional_arrays_1480

func runningSum(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	for i:=1;i<len(nums);i++ {
		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}