package Number_of_islands_200

func numIslands(grid [][]byte) int {
	count := 0
	for line := range grid {
		for column := range grid[line] {
			if grid[line][column] == '0' || grid[line][column] == '2' {
				continue
			}
			count ++
			dfc(grid, line, column)
		}
	}
	return count
}

func dfc(grid [][]byte, x, y int) {
	if x<0||x>=len(grid) ||y<0||y>=len(grid[x]) {
		return
	}
	if grid[x][y] != '1' {
		return
	}
	grid[x][y] = '2'
	dfc(grid, x+1, y)
	dfc(grid, x-1, y)
	dfc(grid, x, y+1)
	dfc(grid, x, y-1)
}