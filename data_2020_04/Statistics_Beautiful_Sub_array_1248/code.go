package Statistics_Beautiful_Sub_array_1248

func numberOfSubarrays(nums []int, k int) int {
	n := len(nums)
	odd := []int{-1}
	ans := 0
	for i, v := range nums {
		if v % 2==1 {
			odd = append(odd, i)
		}
	}
	odd = append(odd, n)
	for i:=1; i<len(odd)-k;i++ {
		ans += (odd[i]-odd[i-1])*(odd[i+k]-odd[i+k-1])
	}
	return ans
}