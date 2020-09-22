package Reverse_string_541


func reverseStr(s string, k int) string {
	sr := []byte(s)
	var i int
	for i<len(sr) {
		l:=i
		r:=i+k-1
		if r >=len(sr) {
			r=len(sr)-1
		}
		for l<r &&l<len(sr) {
			sr[l], sr[r] =sr[r], sr[l]
			l++
			r--
		}
		i=i+2*k
	}
	return string(sr)
}