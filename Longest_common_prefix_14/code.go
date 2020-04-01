package Longest_common_prefix_14

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0{
		return ""
	}
	res := ""
	for i:=0;i<len(strs[0]);i++{
		for j:=1;j<len(strs);j++{
			if len(strs[j])<i || strs[j][i] != strs[0][i] {
				return res
			}
		}
		res += strs[0][:i+1]
	}
	return res
}