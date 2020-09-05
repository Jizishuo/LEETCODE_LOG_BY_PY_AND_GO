package Gemstones_and_stones_771

func numJewelsInStones(J string, S string) int {
	var res int
	for _,v := range J {
		for _, v1 := range S {
			if rune(v)==rune(v1) {
				res ++
			}
		}
	}
	return res
}