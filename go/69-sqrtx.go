func mySqrt(x int) int {
	left := 0
	right := x
	for left < right {
		m := left + (right-left+1)/2
		if m*m > x {
			right = m - 1
		} else {
			left = m
		}
	}
	return left
}