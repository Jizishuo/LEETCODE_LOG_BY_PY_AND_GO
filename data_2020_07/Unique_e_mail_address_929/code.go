package Unique_e_mail_address_929

import (
	"strings"
)

func numUniqueEmails(emails []string) int {
	tmp := []string{}
	hash := make(map[string]int)
	ans := 0
	for i:=0;i<len(emails);i++ {
		tmp = strings.Split(emails[i], "@")
		for j:=0;j<len(tmp[0]);j++ {
			if tmp[0][j] == '.' {
				tmp[0] = tmp[0][:j] + tmp[0][j+1:]
				continue
			}
			if tmp[0][j] == '+' {
				tmp[0] = tmp[0][:j]
				break
			}
		}
		if _, ok:= hash[tmp[0] + "@" + tmp[1]]; !ok {
			hash[tmp[0]+ "@" + tmp[1]] = 1
			ans ++
		}
	}
	return ans
}