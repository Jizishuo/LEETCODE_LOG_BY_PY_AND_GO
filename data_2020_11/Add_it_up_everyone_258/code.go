package Add_it_up_everyone_258

func addDigits(num int) int {
	if num == 0 {
		return 0
	}
	if num % 9 ==0 {
		return 9
	} else {
		return num % 9
	}
}