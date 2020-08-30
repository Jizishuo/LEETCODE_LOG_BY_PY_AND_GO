package Split_the_balance_string_1221

func balancedStringSplit(s string) int {
	res := 0
	b := 0
	for i:=0;i<len(s);i++ {
		if s[i] == 'L' {
			b++
		} else if s[i]=='R' {
			b--
		}

		if b ==0 {
			res ++
		}
	}
	return res
}