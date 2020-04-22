package Roman_numerals_to_integers_13

func romanToInt(s string) int {
	res := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	ii := res[s[len(s)-1]]
	for index:=len(s)-2;index>=0;index-- {
		if res[s[index]]<res[s[index+1]] {
			ii -= res[s[index]]
		} else {
			ii += res[s[index]]
		}
	}
	return ii
}