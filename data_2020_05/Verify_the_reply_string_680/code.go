package Verify_the_reply_string_680

func validPalindrome(s string) bool {
	return search(s, 1)
}

func search(s string, time int) bool {
	begin :=0
	end := len(s)-1
	for begin < end {
		if s[begin] != s[end] {
			if time == 0 {
				return false
			}
			return search(s[begin:end], time-1) || search(s[begin+1:end+1], time-1)

		}
		begin ++
		end --
	}
	return true
}