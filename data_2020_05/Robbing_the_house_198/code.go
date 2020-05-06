package Robbing_the_house_198

func rob(nums []int) int {
	p, c := 0, 0
	for _, v := range nums {
		p, c = c, func(x, y int) int {
			if x>y {
				return x
			} else {
				return y
			}
		}(c, p+v)
	}
	return c
}