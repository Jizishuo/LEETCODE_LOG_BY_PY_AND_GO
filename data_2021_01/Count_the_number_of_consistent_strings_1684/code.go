package Count_the_number_of_consistent_strings_1684

func countConsistentStrings(allowed string, words []string) int {
	res := len(words)
	sm := [26]bool{}
	for i:=0;i<len(allowed);i++ {
		sm[allowed[i]-'a'] = true
	}
	for _, w := range words {
		for _, a := range w {
			if !sm[a-'a'] {
				res --
				break
			}
		}
	}
	return res
}
