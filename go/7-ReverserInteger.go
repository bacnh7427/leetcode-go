func reverse(x int) int {
	result := 0
	isNegative := x < 0
	if isNegative {
		x = x * -1
	}
	for x > 0 {
		r := x % 10
		x = x / 10
		result = result*10 + r
		if result > math.MaxInt32 {
			return 0
		}
	}
	if isNegative {
		result = result * -1
	}
	return result
}