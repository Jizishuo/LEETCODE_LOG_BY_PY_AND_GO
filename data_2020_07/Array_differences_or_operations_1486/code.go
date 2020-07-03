package Array_differences_or_operations_1486

func xorOperation(n int, start int) int {
	res := 0
	for i:=0;i<n;i++{
		res = res^start + 2*i
	}
	return res
}