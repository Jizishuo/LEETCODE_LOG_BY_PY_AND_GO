package Replace_spaces_05

import "strings"

func replaceSpace(s string) string {
	//ss := []byte(s)
	//res := []byte{}
	//for i:=0;i<len(ss);i++{
	//	if ss[i] == ' ' {
	//		res = append(res, []byte("%20")...)
	//	} else {
	//		res = append(res, ss[i])
	//	}
	//}
	//return string(res)
	//var b []rune
	//for _,v :=range s{
	//	if v==32{
	//		b =append(b,37,50,48)
	//	}else{
	//		b =append(b,v)
	//	}
	//}
	//return string(b)
	return strings.ReplaceAll(s, " ", "%20")
}