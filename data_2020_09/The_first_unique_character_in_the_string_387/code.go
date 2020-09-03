package The_first_unique_character_in_the_string_387

func firstUniqChar(s string) int {
	var aa [26]int
	for i,k := range s {
		aa[k - 'a'] = i
	}
	for i,k := range s {
		if i ==aa[k-'a'] {
			return i
		} else {
			aa[k-'a'] = -1
		}
	}
	return -1
}