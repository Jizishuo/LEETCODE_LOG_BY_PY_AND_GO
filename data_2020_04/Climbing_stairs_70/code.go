package Climbing_stairs_70

func climbStairs(n int) int {
	one := 1
	twe := 2
	res := 0
	for i:=2; i <n;i++ {
		res = one+twe

		one = twe
		twe = res
	}
	return func(x, y int) int {
				if x>y {
					return x
				} else {
					return y
				}
			}(n, res)

}