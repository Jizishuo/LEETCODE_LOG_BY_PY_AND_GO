package Unique_number_of_appearances_1207

func uniqueOccurrences(arr []int) bool {
	m := map[int]int{}
	for _, v := range arr {
		m[v] ++
	}
	b := map[int]bool{}
	for _, v := range m {
		if b[v] {
			return false
		} else {
			b[v] = true
		}
	}
	return true
}