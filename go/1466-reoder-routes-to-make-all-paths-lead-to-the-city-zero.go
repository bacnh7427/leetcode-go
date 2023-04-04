func minReorder(n int, connections [][]int) int {

	m := make(map[int][]int)
	st := make(map[string]bool)
	for _, v := range connections {
		str := strconv.Itoa(v[0]) + "," + strconv.Itoa(v[1])
		st[str] = true
		m[v[0]] = append(m[v[0]], v[1])
		m[v[1]] = append(m[v[1]], v[0])

	}
	visited := make(map[int]bool)
	var count int
	dfs(0, m, visited, st, &count)
	return count
}

func dfs(node int, m map[int][]int, visited map[int]bool, st map[string]bool, count *int) {
	visited[node] = true
	for _, v := range m[node] {
		if !visited[v] {
			str := strconv.Itoa(node) + "," + strconv.Itoa(v)
			if st[str] {
				*count++
			}
			dfs(v, m, visited, st, count)
		}
	}
}