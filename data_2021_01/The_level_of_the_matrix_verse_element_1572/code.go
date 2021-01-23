package The_level_of_the_matrix_verse_element_1572

func diagonalSum(mat [][]int) int {
	l, r := len(mat), 0

	for i:=0;i<l;i++ {
		r += mat[i][i]
		mat[i][i] = 0
		r += mat[i][l-i-1]
	}
	return r
}