package Word_law_290

import (
	"strings"
)

func wordPattern(pattern string, str string) bool {
	strarr := strings.Split(str, " ")
	if len(strarr) != len(pattern) {
		return false
	}
	h := map[byte]string{}
	h2 := map[string]byte{}
	for i:=0;i<len(pattern);i++{
		v, ok := h[pattern[i]]
		v2 ,ok2 := h2[strarr[i]]
		if ok && v!=strarr[i] || ok2 && v2 !=pattern[i] {
			return false
		} else {
			h[pattern[i]] = strarr[i]
			h2[strarr[i]] = pattern[i]
		}
	}
	return true
}