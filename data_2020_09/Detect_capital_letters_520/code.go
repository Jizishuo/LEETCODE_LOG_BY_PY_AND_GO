package Detect_capital_letters_520

func detectCapitalUse(word string) bool {
	flag := 0
	for i:=0;i<len(word);i++ {
		if word[i]>=65 && word[i]<=90 {
			flag++
		}
	}

	if flag == 0 {
		return true
	} else if flag ==1 && len(word)!=1 && word[0]<=90 {
	 	return true
	} else if flag == len(word) {
		return true
	}
	return false
}