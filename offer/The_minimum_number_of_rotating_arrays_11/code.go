package The_minimum_number_of_rotating_arrays_11

func minArray(numbers []int) int {
	l, r := 0, len(numbers)-1

	for l<=r {
		m := (l+r)/2
		if numbers[m]>numbers[r] {
			l = m+1
		} else {
			r --
		}
	}
	return numbers[l]

}