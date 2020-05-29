package Effective_letter_ectopic_words_242


func isAnagram(s string, t string) bool {
	s_m := [26]int{}
	t_m := [26]int{}

	for _, v := range s {
		s_m[v-'a'] +=1
	}

	for _, v := range t {
		t_m[v-'a'] +=1
	}
	return s_m == t_m
}