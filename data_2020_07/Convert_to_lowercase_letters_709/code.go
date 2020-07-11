package Convert_to_lowercase_letters_709

func toLowerCase(str string) string {
	mm := map[string]string{
		"A": "a", "B": "b", "C": "c", "D": "d", "E": "e", "F": "f",
		"G": "g", "H": "h", "I": "i", "J": "j", "K": "k", "L": "l",
		"M": "m", "N": "n", "O": "o", "P": "p", "Q": "q", "R": "r",
		"S": "s", "T": "t", "U": "u", "V": "v", "W": "w", "X": "x",
		"Y": "y", "Z": "z"}

	ret := make([]byte, 0)
	for _,x := range str {
		if vv, ok := mm[string(x)];ok {
			ret = append(ret, vv...)
		} else {
			ret = append(ret, string(x)...)
		}
	}
	return string(ret)
}