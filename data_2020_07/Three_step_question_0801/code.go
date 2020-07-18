package Three_step_question_0801

func waysToStep(n int) int {
	a,b,c := 4, 2, 1
	if n < 3{
		return n
	}
	if n == 3 {
		return n
	}
	for i:=3;i<n;i++ {
		a,b,c = (a+b+c)%1000000007,a,b
	}
	return a
}