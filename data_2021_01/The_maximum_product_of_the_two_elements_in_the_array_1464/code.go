package The_maximum_product_of_the_two_elements_in_the_array_1464

func maxProduct(nums []int) int {
	k, max := 0, 0
	for k <2 {
		max = nums[k]
		for i:=k;i<len(nums);i++ {
			if max <nums[i] {
				max = nums[i]
				nums[k], nums[i] = nums[i], nums[k]
			}
		}
		k++
	}
	return (nums[0]-1)*(nums[1]-1)
}
