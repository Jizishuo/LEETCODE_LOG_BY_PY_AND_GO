package Whether_the_robot_can_return_to_its_origin_657

func judgeCircle(moves string) bool {
	a, b := 0, 0
	for i:=0;i<len(moves);i++ {
		switch  moves[i] {
		case 'U':
			b--
		case 'D':
			b++
		case 'L':
			a--
		case 'R':
			a++
		}
	}
	return a==0 && b==0
}