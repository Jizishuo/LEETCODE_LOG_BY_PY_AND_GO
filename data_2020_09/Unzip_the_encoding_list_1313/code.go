package Unzip_the_encoding_list_1313

func decompressRLElist(nums []int) []int {
	res := []int{}
	for i:=1;i<=len(nums)/2;i++ {
		for j:=nums[i*2-2];j>0;j-- {
			res = append(res, nums[i*2-1])
		}
	}
	return res
}