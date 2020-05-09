package Plus_one_66

func plusOne(digits []int) []int {
	ll := len(digits)
	if ll == 0 {
		return []int{1}
	}

	for i:= ll-1;i>=0;i-- {
		if digits[i] != 9 {
			digits[i] ++
			return digits
		} else {
			digits[i] = 0
		}
	}
	return append([]int{1}, digits...)

}