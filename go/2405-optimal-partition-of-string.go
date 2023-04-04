func partitionString(s string) int {
	res := 0
	cnt := 1 << 26
	for i := 0; i < len(s); i++ {
		if cnt&(1<<(s[i]-'a')) > 0 {
			res++
			cnt = 1 << 26
		}
		cnt = cnt | (1 << (s[i] - 'a'))
	}
	return res + 1
}