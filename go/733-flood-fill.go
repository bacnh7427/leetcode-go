func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	solve(image, sr, sc, image[sr][sc], color)
	return image
}

func solve(inp [][]int, sr int, sc int, prevColor int, color int) {
	if sr < 0 || sc < 0 || sr >= len(inp) || sc >= len(inp[0]) || inp[sr][sc] != prevColor || prevColor == color {
		return
	}
	inp[sr][sc] = color
	solve(inp, sr+1, sc, prevColor, color)
	solve(inp, sr-1, sc, prevColor, color)
	solve(inp, sr, sc+1, prevColor, color)
	solve(inp, sr, sc-1, prevColor, color)
}