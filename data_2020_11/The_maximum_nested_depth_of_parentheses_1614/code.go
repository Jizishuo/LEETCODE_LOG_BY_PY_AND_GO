package The_maximum_nested_depth_of_parentheses_1614

func maxDepth(s string) int {
	//m_x := 0
	//t := 0
	//
	//for i:=0;i<len(s);i++ {
	//	if s[i] =='(' {
	//		t++
	//		if m_x < t{
	//			m_x=t
	//		}
	//	}
	//	if s[i]==')' {
	//		t--
	//	}
	//}
	//return m_x
	var xx, c int

	for i:=0;i<len(s);i++ {
		switch b :=s[i];b {
		case '(':c++
		case ')':c--
		}
		if c>xx {
			xx=c
		}
	}
	return xx
}