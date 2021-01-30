package Count_the_triples_1534

func countGoodTriplets(arr []int, a int, b int, c int) int {
	l := len(arr)
	res := 0
	for i:=0;i<l;i++ {
		for j:=i+1;j<l;j++ {
			for k:=j+1;k<l;k++ {
				if abs(arr[i],arr[j]) <= a && abs(arr[j],arr[k]) <= b && abs(arr[i], arr[k]) <= c {
					res ++
				}
			}
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