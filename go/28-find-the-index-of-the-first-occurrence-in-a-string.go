func strStr(haystack string, needle string) int {
	if len(needle) == 0 {
		return 0
	}

	if len(haystack) == 0 {
		return -1
	}

	for i := 0; i <= len(haystack)-len(needle); i++ {
		if haystack[i] == needle[0] {
			if len(needle) == 1 {
				return i
			}

			for j := 1; j < len(needle) && i+j < len(haystack) && haystack[i+j] == needle[j]; j++ {
				if j == len(needle)-1 {
					return i
				}
			}
		}
	}
	return -1
}