package There_is_duplicate_element_219

func containsNearbyDuplicate(nums []int, k int) bool {
	var m = make(map[int]int)

	for i, v := range nums {
		if _, ok := m[v]; ok {
			if i-m[v] <= k {
				return true
			}
		}
		m[v] = i
	}
	return false
}