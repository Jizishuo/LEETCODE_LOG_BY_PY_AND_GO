package Massager_17_16

func massage(nums []int) int {
	now:= 0
	last := 0
	for _, i := range nums {
		last, now = now, func(x, y int) int {
								if x>y {
									return x
								} else {
									return y
								}
							}(last+i, now)

	}
	return now
}