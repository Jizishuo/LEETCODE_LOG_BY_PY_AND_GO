package String_compression_01_06

import (
	"strings"
	"fmt"
)

func compressString(S string) string {
	if len(S)==0{
		return S
	}
	pre := S[0]
	count := 1
	var res strings.Builder
	for i:=1;i<len(S);i++ {
		if pre == S[i] {
			count ++
		} else {
			res.WriteByte(pre)
			res.WriteString(fmt.Sprintf("%d", count))
			pre = S[i]
			count = 1
		}
	}
	res.WriteByte(pre)
	res.WriteString(fmt.Sprintf("%d",count))
	str := res.String()
	if len(str) >= len(S) {
		return S
	}
	return str
}
