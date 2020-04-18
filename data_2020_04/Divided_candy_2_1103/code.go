package Divided_candy_2_1103

func distributeCandies(candies int, num_people int) []int {
	j := 0
	a := make([]int, num_people)
	for candies>0 {
		for i:=0; i<num_people; i++ {
			if candies>j+1 {
				a[i]+= j+1
				candies -=j+1
			} else {
				a[i] += candies
				candies=0
			}
			j ++
		}

	}
	return a
}