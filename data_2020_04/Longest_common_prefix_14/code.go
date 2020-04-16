package Longest_common_prefix_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefixMap := make(map[int]string)

	firstStr := strs[0]
	for index := 0; index < len(firstStr); index++ {
		s := string(firstStr[index])
		prefixMap[index] = s
	}

	prefixNum := len(firstStr)
	for i := 1; i < len(strs); i++ {
		str := strs[i]
		num := 0
		for j := 0; j < len(str); j++ {
			s := string(str[j])
			if prefixMap[j] == s {
				num++
			}else{
				break
			}
		}
		if num == 0{
			return ""
		}
		if num < prefixNum {
			prefixNum = num
		}
	}
	retStr := ""
	for i := 0;  i < prefixNum; i++ {
		retStr += prefixMap[i]
	}
	return retStr
}