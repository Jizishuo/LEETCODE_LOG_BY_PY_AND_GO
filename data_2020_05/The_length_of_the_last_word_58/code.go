package The_length_of_the_last_word_58

func lengthOfLastWord(s string) int {
	if len(s) == 0 {return 0}
	count := len(s)
	out := 0
	for i := count - 1;i>=0;i--{
		if string(s[i]) != " "{
			out++
		}
		if out != 0 && string(s[i]) == " "{
			break
		}
	}
	return out
}
