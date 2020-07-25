package Subdomain_access_count_811

import (
	"strings"
	"strconv"
	"fmt"
)

func subdomainVisits(cpdomains []string) []string {
	m := make(map[string]int, 0)
	for _, v := range cpdomains {
		varr := strings.Split(v, " ")
		num, _ := strconv.Atoi(varr[0])
		docpd := strings.Split(varr[1], ".")
		for i:=0;i<len(docpd);i++ {
			m[strings.Join(docpd[i:],".")] +=num
		}
	}
	var res []string
	for k, v := range m {
		res = append(res, fmt.Sprintf("%d %s", v,k))
	}
	return res
}