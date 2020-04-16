package Candy_575

func distributeCandies(candies []int) int {
	m := make(map[int]int)
	sum := len(candies)
	for _, c := range candies {
		m[c] += 1
	}
	if len(m) < sum/2 {
		return len(m)
	} else {
		return sum/2
	}
}