package The_difference_between_the_accumulations_of_the_integers_1281

func subtractProductAndSum(n int) int {
	a, b := 0, 1
	for n>0 {
		d := n%10
		n = n/10
		a=a+d
		b=b*d
	}
	return b-a
}