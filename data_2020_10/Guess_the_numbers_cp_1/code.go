package Guess_the_numbers_cp_1

func game(guess []int, answer []int) int {
	res := 0
	for i:=0;i<len(guess);i++ {
		if guess[i]==answer[i] {
			res ++
		}
	}
	return res
}