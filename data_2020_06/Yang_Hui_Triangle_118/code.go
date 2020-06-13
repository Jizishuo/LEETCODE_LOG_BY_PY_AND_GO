package Yang_Hui_Triangle_118

func generate(numRows int) [][]int {
	//res := [][]int{}
	//if numRows == 0 {
	//	return res
	//}
	//res = append(res, []int{1})
	//for i:=1;i<numRows;i++ {
	//	in := []int{0}
	//	in = append(in, res[i-1]...)
	//	for ii :=0;ii<len(in)-1;ii++ {
	//		in[ii] = in[ii]+in[ii+1]
	//	}
	//	res = append(res, in)
	//}
	//return res
	if numRows == 0 {
		return nil
	}
	res := make([][]int, numRows)
	res[0] = []int{1}
	for i := 1; i < numRows; i++ {
		res[i] = make([]int, i+1)
		res[i][0] = 1
		res[i][i] = 1

		for j := 1; j < i; j++ {
			res[i][j] = res[i-1][j-1] + res[i-1][j]
		}
	}
	return res
}