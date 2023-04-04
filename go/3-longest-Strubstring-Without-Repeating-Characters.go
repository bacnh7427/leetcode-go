func lengthOfLongestSubstring(s string) int {
	newMap := make(map[byte]int)
	res := 0
	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := newMap[s[r]]; ok {
			l = max(l, index)
		}
		res = max(res, r-l+1)
		newMap[s[r]] = r + 1
	}
	return res
}

func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}