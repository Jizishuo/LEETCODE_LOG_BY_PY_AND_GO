package The_number_of_operations_to_turn_a_number_into_0_1342

func numberOfSteps(num int) (ans int) {
	for ;; ans++ {
		if num==0{
			break
		} else if num & 1 ==1 {
			num--
		} else {
			num >>=1
		}
	}
	return
}