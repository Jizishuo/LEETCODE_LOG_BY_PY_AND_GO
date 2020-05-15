package Determine_whether_a_character_is_unique_0101

func isUnique(astr string) bool {
	m := make([]byte, 128)
	for s := range astr {
		if m[astr[s]] > 0 {
			return false
		}
		m[astr[s]] ++
	}
	return true
}