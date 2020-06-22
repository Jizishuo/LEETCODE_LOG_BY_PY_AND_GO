package Regular_expression_match_19

func isMatch(s string, p string) bool {
	if len(p)==0{
		return len(s)==0
	}
	var ff bool
	if len(s)>0 && (s[0]==p[0] || p[0]=='.') {
		ff = true
	}
	if len(p)>1 && p[1]== '*' {
		return (ff && isMatch(s[1:], p) || isMatch(s, p[2:]))
	}
	return ff && isMatch(s[1:], p[1:])
}