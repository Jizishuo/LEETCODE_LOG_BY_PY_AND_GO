package Left_rotation_string_58

func reverseLeftWords(s string, n int) string {
	//res := append([]byte(s[n:]), []byte(s[:n])...)
	//return string(res)
	return s[n:] + s[:n]
}