package Happy_Numbers_202

func isHappy(n int) bool {
	h := make(map[int]bool)
	for n != 1 {
		if _, ok := h[n];ok {
			return false
		}
		h[n] = true
		var res int
		for n != 0 {
			res += (n%10)*(n%10)
			n = n/10
		}
		n = res
	}
	return true
}