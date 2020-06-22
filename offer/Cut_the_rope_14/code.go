package Cut_the_rope_14

func cuttingRope(n int) int {
	if n == 2 {
		return 1
	}
	if n == 3 {
		return 2
	}
	dp:= make([]int, n+1)
	dp[2] = 2
	dp[3] = 3

	for i:=4;i<n+1;i++ {
		ii := i-1
		j := 1
		ans := 0
		for ii>=j {
			if dp[ii]*j >ans {
				ans = dp[ii]*j
			}
			ii--
			j++
		}
		dp[i]=ans
	}
	return dp[n]
}