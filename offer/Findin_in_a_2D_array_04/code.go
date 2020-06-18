package Findin_in_a_2D_array_04

func findNumberIn2DArray(matrix [][]int, target int) bool {
	i, j := len(matrix)-1, 0
	l := len(matrix[0])
	for i>=0 && j<l {

		if matrix[i][j]> target{
			i --
		} else if matrix[i][j] < target {
			j ++
		} else {
			return true
		}
	}
	return false
}