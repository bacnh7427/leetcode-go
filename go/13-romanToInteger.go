func romanToInt(s string) int {
	var result int
	var current int
	var prev int
	for i := 0; i < len(s); i++ {
		current = symbolValue(s[i])
		result += current
		if current > prev {
			result = result - prev*2
		}
		prev = current
	}
	return result
}

func symbolValue(symbol byte) int {
	switch symbol {
	case 'I':
		return 1
	case 'V':
		return 5
	case 'X':
		return 10
	case 'L':
		return 50
	case 'C':
		return 100
	case 'D':
		return 500
	case 'M':
		return 1000
	}
	return 0
}