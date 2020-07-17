package An_effective_array_of_mountains_941

func validMountainArray(A []int) bool {
	//up, down := false, false
	var (
		up  bool
		down bool
	)

	for i:=1;i<len(A);i++ {
		if A[i]-A[i]>0 {
			if down {
				return false
			}
			up = true
		} else if A[i]-A[i]<0 {
			if !up {
				return false
			}
			down = true
		} else {
			return false
		}
	}
	return up && down
}