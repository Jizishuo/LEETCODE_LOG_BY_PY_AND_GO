package Verify_the_reply_string_125

func isPalindrome(s string) bool {
	var arr []rune
	for _, ch := range s {
		if ch >='a' && ch <='z'|| ch >='0' && ch <='9' {
			arr = append(arr, ch)
		} else if ch >= 'A' && ch <= 'Z' {
			arr = append(arr, ch+('a'-'A'))
		} else {
			continue
		}
	}
	for i :=0;i<len(arr)/2;i++ {
		if arr[i]!=arr[len(arr)-i-1] {
			return false
		}
	}
	return true
}