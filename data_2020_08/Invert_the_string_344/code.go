package Invert_the_string_344

func reverseString(s []byte)  {
	i, j := 0, len(s)-1
	for i <=j {
		s[i], s[j] = s[j], s[i]
	}
}