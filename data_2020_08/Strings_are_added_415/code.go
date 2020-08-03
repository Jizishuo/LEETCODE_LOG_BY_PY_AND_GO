package Strings_are_added_415

import (
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	c := 0
	res := ""
	for i,j := len(num1)-1, len(num2)-1;i>=0 ||j>=0|| c !=0;i,j=i-1,j-1 {
		var n1, n2 int
		if i >=0 {
			n1 = int(num1[i]-'0')
		}
		if j >=0 {
			n2 = int(num2[j]-'0')
		}
		tmp := n1+n2 +c

		res = strconv.Itoa(tmp%10)+res
		c = tmp /10
	}
	return res

}