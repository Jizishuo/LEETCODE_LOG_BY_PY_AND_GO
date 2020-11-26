package Hanming_distance_461

func hammingDistance(x int, y int) int {
	i := x^y
	c := 0
	for i!=0 {
		if i&1==1 {
			c++
		}
		i = i>>1
	}
	return c
}