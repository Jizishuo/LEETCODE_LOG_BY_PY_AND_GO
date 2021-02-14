package The_distance_value_between_the_two_arrays_1385

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res :=0
	for i:=0;i<len(arr1);i++ {
		isok := true
		for ii:=0;ii<len(arr2);ii++ {
			if abs(arr1[i], arr2[ii]) <= d {
				isok = false
			}
		}
		if isok {
			res ++
		}

	}
	return res
}

func abs(a, b int) int {
	if a>b {
		return a-b
	} else {
		return b-a
	}
}