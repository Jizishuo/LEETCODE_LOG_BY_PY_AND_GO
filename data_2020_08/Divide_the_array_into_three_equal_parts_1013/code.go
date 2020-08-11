package Divide_the_array_into_three_equal_parts_1013

func canThreePartsEqualSum(A []int) bool {
	var sum int
	for _, v := range A {
		sum = v + sum
	}
	if sum%3 != 0 {
		return false
	}
	num, count := 0, 0
	for k,v := range A {
		num += v
		if num == sum/3 {
			num = 0
			count ++
			if count==2 && k != len(A)-1 {
				return true
			}
		}
	}
	return false

}