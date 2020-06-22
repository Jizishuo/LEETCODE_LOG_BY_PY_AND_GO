package A_string_representing_a_value_20

import (
	"strings"
	"regexp"
)

func isNumber(s string) bool {
	s = strings.TrimSpace(s)
	res,_ := regexp.MatchString("^(([\\+\\-]?[0-9]+(\\.[0-9]*)?)|([\\+\\-]?\\.?[0-9]+))(e[\\+\\-]?[0-9]+)?$", s)

	return res
}
