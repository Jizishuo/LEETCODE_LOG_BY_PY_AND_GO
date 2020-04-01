package palindromic_number_9

func isPalindrome(x int) bool {
	s := fmt.Sprintf("%v", x)
	if len(s)<=1 {
		return true
	}
	i:=0
	j:=len(s)-1
	for i<j {
		if s[i] != s[j]{
			return false
		}
		i++
		j--
	}
	return true
}