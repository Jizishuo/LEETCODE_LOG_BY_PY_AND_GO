package Yang_Hui_Triangle_119

func getRow(rowIndex int) []int {
	res := []int{1}
	for i:=1;i<=rowIndex;i++ {
		res = append(res, 1)
		for j:=i-1;j>0;j-- {
			res[j]+=res[j-1]
		}
	}
	return res
}