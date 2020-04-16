package Flip_words_in_a_string_151

import "strings"

func reverseWords(s string) string {
	var reserseSeg []string

	seg := strings.Fields(s)
	for i := len(seg)-1; i>=0; i-- {
		reserseSeg =append(reserseSeg, seg[i])
	}
	return strings.Join(reserseSeg, " ")
}