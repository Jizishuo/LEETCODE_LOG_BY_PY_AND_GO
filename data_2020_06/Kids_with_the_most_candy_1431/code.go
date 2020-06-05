package Kids_with_the_most_candy_1431

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxn := -1
	l := len(candies)
	for i :=0;i<l;i++{
		if candies[i]>maxn {
			maxn = candies[i]
		}
	}

	res := []bool{}
	for _, n := range candies {
		if n+extraCandies >= maxn {
			res = append(res, true)
		} else {
			res = append(res, false)
		}
	}
	return res
}