package Reverse_the_word_in_the_string_557

func reverseWords(s string) string {
	b := []byte(s)
	l := 0
	for i,v := range s {
		if v == ' ' || i == len(s)-1 {
			r := i-1
			if i==len(s)-1 {
				r = i
			}
			for l<r {
				b[l],b[r]=b[r], b[l]
				l++
				r--
			}
			l =i+1
		}
	}
	return string(b)
}