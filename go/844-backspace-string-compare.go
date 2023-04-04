func backspaceCompare(s string, t string) bool {
	stackS := buildStack(s)
	stackT := buildStack(t)
	return string(stackS) == string(stackT)
}

func buildStack(a string) []byte {
	var result []byte
	for i := 0; i < len(a); i++ {
		if a[i] == '#' {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, a[i])
		}
	}
	return result
}