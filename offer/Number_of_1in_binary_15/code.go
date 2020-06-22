package Number_of_1in_binary_15

func hammingWeight(num uint32) int {
	res := 0
	for num>0 {
		if num & 1 ==1{
			res ++
		}
		num >>= 1
	}
	return res
}