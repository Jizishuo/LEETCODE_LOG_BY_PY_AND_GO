package The_integer_sub_side_of_the_value_16

func myPow(x float64, n int) float64 {
	if n < 0 {
		x = 1.0/x
		n = -n
	} else if n == 0 {
		return 1
	}
	res := 1.0
	for n >1 {
		if n &1 == 1 {
			res *= x
			n--
		} else {
			x = x*x
			n = n/2
		}
	}
	res *= x
	return res
}