func gcdOfStrings(str1 string, str2 string) string {
	len1 := len(str1)
	len2 := len(str2)
	maxLen := max(len1, len2)
	for i := maxLen; i >= 1; i-- {
		if len1%i == 0 && len2%i == 0 && str1[0:i] == str2[0:i] {
			tmp1 := str1[i:] + str1[0:i]
			tmp2 := str2[i:] + str2[0:i]
			if tmp1 == str1 && tmp2 == str2 {
				return str1[0:i]
			}
		}
	}

	return ""
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
