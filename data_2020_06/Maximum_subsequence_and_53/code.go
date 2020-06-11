package Maximum_subsequence_and_53

func maxSubArray(nums []int) int {

	maxx := nums[0]
	for i:=1; i<len(nums); i++ {
		if nums[i] + nums[i-1] > nums[i] {
			nums[i] += nums[i-1]
		}
		if nums[i] > maxx {
			maxx = nums[i]
		}
	}
	return maxx
}