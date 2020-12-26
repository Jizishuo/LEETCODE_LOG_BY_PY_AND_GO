package Rise_and_fall_strings_1370

func sortString(s string) string {
	cut := ['z'+1]int{}
	for _, ch := range s {
		cut[ch] ++
	}
	n := len(s)
	ans := make([]byte, 0, n)
	for len(ans) <n {
		for i:= byte('a');i<='z';i++ {
			if cut[i] > 0 {
				ans = append(ans, i)
				cut[i] --
			}
		}

		for i:=byte('z');i>='a';i-- {
			if cut[i]>0 {
				ans = append(ans, i)
				cut[i]--
			}
		}

	}
	return string(ans)
}