package Passing_information_lcp_07

func numWays(n int, relation [][]int, k int) int {
	d := make([]int, n)
	d[0] = 1
	for i:=0;i<k;i++ {
		t := make([]int, n)
		for _, j := range relation {
			t[j[1]] += d[j[0]]
		}
		d = t
	}
	return d[n-1]
}