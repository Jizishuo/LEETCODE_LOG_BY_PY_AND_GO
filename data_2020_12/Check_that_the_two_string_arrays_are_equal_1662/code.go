package Check_that_the_two_string_arrays_are_equal_1662

import (
	"strings"
)

func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}