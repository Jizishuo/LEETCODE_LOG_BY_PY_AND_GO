package Paths_in_the_matrix_12

func exist(board [][]byte, word string) bool {

	for i:=0;i<len(board);i++ {
		for j:=0;j<len(board[0]);j++ {

			if sss(board, word,i,j,0) {
				return true
			}
		}
	}
	return false
}


func sss(board [][]byte, word string, i,j,k int) bool {
	if k == len(word) {
		return true
	}
	if i<0 ||j<0 ||i == len(board) || j == len(board[0]) {
		return false
	}
	if board[i][j] == word[k] {
		tmp := board[i][j]
		board[i][j] = ' '
		if sss(board, word, i,j+1,k+1) ||
			sss(board, word, i, j-1, k+1) ||
			sss(board, word, i+1, j, k+1) ||
			sss(board, word, i-1, j, k+1) {
			return true
		} else {
			board[i][j] = tmp
		}
	}
	return false
}