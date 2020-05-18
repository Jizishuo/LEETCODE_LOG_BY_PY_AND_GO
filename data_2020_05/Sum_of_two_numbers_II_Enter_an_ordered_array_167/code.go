package Sum_of_two_numbers_II_Enter_an_ordered_array_167

func twoSum(numbers []int, target int) []int {
	//for i:=0;i<len(numbers);i++ {
	//
	//	for j:=i+1;j<len(numbers);j++ {
	//
	//		if numbers[i]+numbers[j]==target {
	//			return []int{i, j}
	//		}
	//	}
	//}
	//return []int{}
	i := 0
	j := len(numbers)-1

	for i<j {
		if numbers[i]+numbers[j] == target {
			return []int{i+1, j+1}
		} else if numbers[i]+numbers[j] <target {
			i++
		} else {
			j --
		}
	}
	return []int{}
}