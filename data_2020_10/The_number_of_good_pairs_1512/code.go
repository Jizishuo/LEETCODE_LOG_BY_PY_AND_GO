package The_number_of_good_pairs_1512

func numIdenticalPairs(nums []int) int {
	res := 0
	for i:=0;i<len(nums)-1;i++ {
		for n:=i+1;n<len(nums);n++ {
			if nums[i]==nums[n] {
				res ++
			}
		}
	}
	return res
}