package Number_of_pieces_that_can_be_captured_in_one_step_999

func numRookCaptures(board [][]byte) int {

	var Ri, Rj int

	for i1, j1 := range board {
		for i2, j2 := range j1 {
			if j2 == 'R' {
				Ri = i1
				Rj = i2
			}
		}
	}

	var count int

	for i := Rj;i<8; i++ {
		if board[Ri][i] == 'p' {
			count++
			break
		}
		if board[Ri][i] == 'B' {
			break
		}
	}
	for i := Rj;i>0; i-- {
		if board[Ri][i] == 'p' {
			count++
			break
		}
		if board[Ri][i] == 'B' {
			break
		}
	}

	for i := Ri;i<8; i++ {
		if board[i][Rj] == 'p' {
			count++
			break
		}
		if board[i][Rj] == 'B' {
			break
		}
	}
	for i := Ri;i>0; i-- {
		if board[i][Rj] == 'p' {
			count++
			break
		}
		if board[i][Rj] == 'B' {
			break
		}
	}

	return count
}