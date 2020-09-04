package Look_for_a_difference_389

func findTheDifference(s string, t string) byte {
	hash := make(map[byte]int)
	var res byte

	for k, _ := range t {
		hash[t[k]]++
	}

	for k, _ := range s {
		hash[s[k]] --
	}

	for k, v := range hash {
		if v ==1 {
			res = k
		}
	}
	return res
}