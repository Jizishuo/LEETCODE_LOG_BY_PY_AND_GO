package Longest_substring_without_repeating_characters_3

func lengthOfLongestSubstring(s string) int {
	var length int
	var s1 string
	left := 0
	right := 0
	s1 =s[left: right]
	for ; right<len(s); right++ {
		if index := strings.IndexByte(s1, s[right]);index != -1 {
			left += index +1
		}
		s1 = s[left: right+1]
		if len(s1)> length {
			length = len(s1)
		}
	}
	return length

}