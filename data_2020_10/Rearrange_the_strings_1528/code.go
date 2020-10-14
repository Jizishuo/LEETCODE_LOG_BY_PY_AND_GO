package Rearrange_the_strings_1528

func restoreString(s string, indices []int) string {
	res :=make([]byte, len(indices))
	for i,v := range indices {
		res[v]=s[i]
	}
	return string(res)
}