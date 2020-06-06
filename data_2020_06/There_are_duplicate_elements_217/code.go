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
}