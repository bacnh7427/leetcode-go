func longestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}

	digitPalind := make([][]int, len(s), len(s))
	for i := range digitPalind {
		digitPalind[i] = make([]int, len(s), len(s))
	}

	var start, length int
	for i := len(s) - 1; i >= 0; i-- {
		for j := i; j < len(s); j++ {
			if s[i] == s[j] {
				if j-i < 3 {
					digitPalind[i][j] = 1

				} else {
					digitPalind[i][j] = digitPalind[i+1][j-1]
				}
			}

			if digitPalind[i][j] != 1 {
				continue
			}
			if j-i+1 > length {
				length = j - i + 1
				start = i
			}
		}
	}
	return string(s[start : start+length])
}