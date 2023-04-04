func compress(chars []byte) int {
	var result int
	for i := 0; i < len(chars); {
		j := i + 1
		for j < len(chars) && chars[i] == chars[j] {
			j++
		}
		chars[result] = chars[i]
		result++
		if j-i > 1 {
			s := fmt.Sprintf("%d", j-i)
			for k := 0; k < len(s); k++ {
				chars[result] = s[k]
				result++
			}
		}
		i = j
	}
	return result
}