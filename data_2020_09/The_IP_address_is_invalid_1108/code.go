package The_IP_address_is_invalid_1108

import "strings"
func defangIPaddr(address string) string {
	res := strings.Builder{}
	for i := range address {
		if address[i]=='.' {
			res.WriteString("[.]")
		} else {
			res.WriteByte(address[i])
		}
	}
	return res.String()
}