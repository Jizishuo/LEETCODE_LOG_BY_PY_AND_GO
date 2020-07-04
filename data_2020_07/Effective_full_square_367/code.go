package Effective_full_square_367

func isPerfectSquare(num int) bool {
	l := 0
	r := num
	for l<=r {
		m := l+ (r-l)/2
		mnum := m*m
		if mnum <num {
			l = m+1
		} else if mnum >num {
			r=m-1
		} else {
			return true
		}
	}
	return false
}