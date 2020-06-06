package There_are_duplicate_elements_217

func containsDuplicate(nums []int) bool {
	//d := map[int]int{}
	//for _, i := range nums {
	//	if _, ok:=d[i];ok {
	//		return true
	//	} else {
	//		d[i] = 1
	//	}
	//
	//}
	//return false
	m := make(map[int]struct{},0)
	for _,v:=range nums{
		if _,ok:=m[v];ok{
			return true
		}
		m[v]= struct{}{}
	}
	return false

	作者：ba-xiang
	链接：https://leetcode-cn.com/problems/contains-duplicate/solution/go-hash-by-ba-xiang-4/
	来源：力扣（LeetCode）
	著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
}