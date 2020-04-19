package Continuous_positive_sequence_57

func findContinuousSequence(target int) [][]int {
	res := make([][]int, 0)
	for i:=int(math.Sqrt(float64(2*target))); i>=2; i-- {
		temp := target-i*(i-1)/2
		if temp%i ==0 {
			begin := temp/i
			mes := make([]int, 0)
			for j:=0;j<i; j++ {
				mes = append(mes, begin+j)
			}
			res = append(res, mes)
		}

	}
	return res
}