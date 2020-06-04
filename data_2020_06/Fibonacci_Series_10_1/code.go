package Fibonacci_Series_10_1

func fid(n int) int {
	//if n < 2{
	//	return n
	//} else {
	//	var(
	//		a = 0
	//		b = 1
	//		t int
	//	)
	//	for i:=2;i<=n;i++{
	//		t =b
	//		b = (a+b)%1000000007
	//		a = t
	//	}
	//	return b
	//}
	if n<2{
		return n
	} else {
		return (fid(n-1)+fid(n-2))%1000000007
	}
}