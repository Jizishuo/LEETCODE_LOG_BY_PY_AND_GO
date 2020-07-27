package Determin_of_the_score_392

func isSubsequence(s string, t string) bool {
	if len(s)==0 {
		return true
	}
	for i:=0;i<len(t);i++ {
		if t[i]==s[0] {
			s = s[1:]
		}
		if len(s)==0 {
			return true
		}
	}
	return false
}