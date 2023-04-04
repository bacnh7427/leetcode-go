func uniquePaths(m int, n int) int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if i == 0 || j == 0 {
				result[i][j] = 1
				continue
			}
			result[i][j] = result[i-1][j] + result[i][j-1]
		}
	}
	return result[n-1][m-1]
}