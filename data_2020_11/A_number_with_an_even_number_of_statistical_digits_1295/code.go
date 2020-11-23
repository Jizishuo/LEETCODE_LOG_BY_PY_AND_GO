package A_number_with_an_even_number_of_statistical_digits_1295

func findNumbers(nums []int) int {
	rr := 0
	for _, v := range nums {
		c:=0
		for 0<v {
			v/=10
			c++
		}
		if c%2==0 {
			rr++
		}
	}
	return rr
}