package The_sum_of_the_squares_633

import "math"

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i<=j {
		sum := i*i + j*j
		if sum ==c {
			return true
		} else if sum <c {
			i++
		} else {
			j--
		}
	}
	return false
}