package Unusual_words_in_two_sentences_884

import (
	"strings"
)

func uncommonFromSentences(A string, B string) []string {
	a := strings.Split(A, " ")
	b := strings.Split(B, " ")
	all := append(a, b...)
	dd := map[string]int{}
	for _, v := range all {
		dd[v] += 1
	}
	//hhh(&dd, A)
	//hhh(&dd, B)
	res := make([]string, 0)
	for k, v := range dd {
		if v == 1 {
			res = append(res, k)
		}
	}
	return res
}
func hhh(m *map[string]int, s string)  {
	sarr := strings.Split(s, " ")
	for _, v := range sarr {
		(*m)[v] += 1
	}
}