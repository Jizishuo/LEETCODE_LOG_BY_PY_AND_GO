package A_number_less_than_the_current_number_1365

func smallerNumbersThanCurrent(nums []int) []int {
	res := make([]int, 0)
	var tmp int
	for index1, v1 := range nums {
		for index2, v2 := range nums {
			if index1 != index2 && v1 >v2 {
				tmp ++
			}
		}
		res = append(res, tmp)
		tmp = 0
	}
	return res
}