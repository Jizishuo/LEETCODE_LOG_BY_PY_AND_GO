package Full_arrangement_46

func permute(nums []int) [][]int {
	if len(nums) <= 1 {
		return [][]int{nums}
	}
	res := [][]int{}
	for i, v := range nums {
		//res_nums := make([]int, len(nums)-1)
		//copy(res_nums[0:], nums[0:i])
		//copy(res_nums[i:], nums[i+1:])
		res_nums := append(nums[:0], nums[i+1])
		for _, vv := range permute(res_nums){
			// res = append([][]int{v}, vv)
			res = append(res, append(vv, v))
		}

	}
	return res
}