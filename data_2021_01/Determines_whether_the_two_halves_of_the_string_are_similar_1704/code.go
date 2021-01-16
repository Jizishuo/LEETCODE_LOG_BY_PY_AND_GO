package Determines_whether_the_two_halves_of_the_string_are_similar_1704

func halvesAreAlike(s string) bool {
	set := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
		'A': true,
		'E': true,
		'I': true,
		'O': true,
		'U': true,
	}
	r := 0
	for i:=0;i<len(s)/2;i++ {
		if _, ok := set[s[i]];ok {
			r ++
		}
		if _, ok := set[s[len(s)-1-i]];ok {
			r --
		}
	}
	return r == 0
}