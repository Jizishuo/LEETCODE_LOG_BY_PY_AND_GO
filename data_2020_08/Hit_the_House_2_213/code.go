package Hit_the_House_2_213

func rob(nums []int) int {
	ll := len(nums)
	if ll == 0 {
		return 0
	}
	if ll <=1 {
		return nums[0]
	}
	dd := func(ns []int, s, n int) int {
		c, p := 0, 0
		for i:=s;i<=n;i++ {
			p, c = c, max(p+ns[i], c)
		}
		return c
	}
	return max(dd(nums, 0,ll-2), dd(nums, 1, ll-1))
}

func max(x, y int) int {
	if x >y {
		return x
	} else {
		return y
	}
}
//func rob(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return nums[0]
//	} else if n == 2 {
//		return Max(nums[0], nums[1])
//	}
//	return Max(startRob(nums, 0, n-2), startRob(nums, 1, n-1))
//}
//// 从start到end开始
//func startRob(nums []int, start, end int) int {
//	pre, max := 0, 0
//	for i := start; i <= end; i++ {
//		pre, max = max, Max(pre+nums[i], max)
//	}
//	return max
//}
//func Max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}
