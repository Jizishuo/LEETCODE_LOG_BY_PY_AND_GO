package Print_from_1t__maximum_n_digits_17

func printNumbers(n int) []int {
	if n == 0 {
		return []int{}
	}
	sum :=1
	for n >0{
		sum =sum*10
		n--
	}
	res := []int{}
	for i:=1;i<sum;i++ {
		res = append(res, i)
	}
	return res
}