func numIslands(grid [][]byte) int {
	countInslands := 0
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' {
				countInslands++
				teraform(i, j, grid)
			}
		}
	}
	return countInslands
}

func teraform(i int, j int, grid [][]byte) {
	if i < 0 || i > len(grid)-1 {
		return
	}

	if j < 0 || j > len(grid[0])-1 {
		return
	}
	if string(grid[i][j]) == "1" {
		grid[i][j] = []byte("0")[0]
		teraform(i+1, j, grid)
		teraform(i-1, j, grid)
		teraform(i, j+1, grid)
		teraform(i, j-1, grid)
	}
}
