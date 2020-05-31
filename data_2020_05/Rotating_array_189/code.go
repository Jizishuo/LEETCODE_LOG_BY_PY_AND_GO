package Rotating_array_189

func rotate(nums []int, k int)  {
	k %=len(nums)
	copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
}