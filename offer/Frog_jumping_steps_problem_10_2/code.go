package Frog_jumping_steps_problem_10_2

func numWays(n int) int {
	a, b := 1, 2
	for i:=1;i<n;i++ {
		a, b = b, (a+b) % 1000000007
	}
	return a
}