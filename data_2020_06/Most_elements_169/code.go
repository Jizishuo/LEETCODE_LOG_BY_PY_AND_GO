package Most_elements_169

func majorityElement(nums []int) int {
	m := 0
	c := 0

	for _, n := range nums {
		if c==0 {
			m = n
		}
		if n == m {
			c ++
		} else {
			c --
		}
	}
	return m
}