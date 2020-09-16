package The_collection_of_errors_645


func findErrorNums(nums []int) []int {
	res1,res2 := 0,0
	tmp := make([]int,len(nums) + 1)
	for i := 0;i < len(nums);i++ {
		tmp[nums[i]]++
	}
	for j := 1;j < len(tmp);j++ {
		if tmp[j] == 2 {
			res1 = j
		}
		if tmp[j] == 0 {
			res2 = j
		}
	}
	return []int{res1,res2}

}