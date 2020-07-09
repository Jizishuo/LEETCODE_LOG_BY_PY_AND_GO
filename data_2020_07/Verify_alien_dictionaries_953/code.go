package Verify_alien_dictionaries_953


func isAlienSorted(words []string, order string) bool {
	m := make(map[int32]int)
	for k, v := range order {
		m[v] = k
	}
	for i:=0;i<len(words)-1;i++ {
		if ccc(words[i], words[i+1], m) {
			return false
		}
	}
	return true
}

func ccc(x, y string, m map[int32]int) bool {
	for i:=0;i<len(x);i++ {
		if i>=len(y) {
			return true
		}
		if m[int32(x[i])] < m[int32(y[i])] {
			return false
		}
		if m[int32(x[i])] > m[int32(y[i])] {
			return true
		}
	}
	return false
}