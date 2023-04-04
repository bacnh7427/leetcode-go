func getHint(secret string, guess string) string {
	b := 0
	c := 0
	tmp := make([]int, 10)
	for i := 0; i < len(secret); i++ {
		s := secret[i] - '0'
		g := guess[i] - '0'
		if s == g {
			b++
			continue
		}
		if tmp[s] < 0 {
			c++
		}

		if tmp[g] > 0 {
			c++
		}

		tmp[s]++
		tmp[g]--
	}
	return fmt.Sprintf("%dA%dB", b, c)
}