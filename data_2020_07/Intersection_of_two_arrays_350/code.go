package Intersection_of_two_arrays_350

func intersect(nums1 []int, nums2 []int) []int {
	m := map[int]int{}
	for _,v := range nums1 {
		if _, ok :=m[v];ok {
			m[v] ++
		} else {
			m[v] = 1
		}
	}
	res := make([]int, 0)
	for _, v := range nums2 {
		if vv, ok :=m[v];ok && vv>0 {
			res = append(res, vv)
			m[v] --
		}
	}
	return res
}