package The_longest_string_of_replies_409

func longestPalindrome(s string) int {
	m := make(map[byte]int)
	for i:=0;i<len(s);i++ {
		m[s[i]] ++
	}
	res := 0
	for _, v := range m {
		res += (v/2)*2
	}
	if res == len(s) {
		return res
	} else {
		return res +1
	}
}