package Find_common_characters_1002

import (
	"math"
)

func commonChars(A []string) []string {
	ans := []string{}
	minFrep := [26]int{}
	for i := range minFrep {
		minFrep[i] = math.MaxInt64
	}

	for _, word := range A {
		frep := [26]int{}
		for _, b := range word {
			frep[b-'a']++
		}
		for i, f := range frep[:] {
			if f < minFrep[i] {
				minFrep[i] = f
			}
		}
	}
	for i:=byte(0);i<26;i++ {
		for j := 0; j < minFrep[i]; j++ {
			ans = append(ans, string('a'+i))
		}
	}
	return ans
}