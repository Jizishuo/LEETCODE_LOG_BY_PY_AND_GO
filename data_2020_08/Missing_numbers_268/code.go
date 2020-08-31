package Missing_numbers_268

func missingNumber(nums []int) int {
	miss := 0

	for i, n := range nums{
		miss ^= i ^ n
	}
	miss ^=len(nums)
	return miss
}