func countOdds(low int, high int) int {
	var count = 0
	if low%2 != 0 {
		count++
	}
	if high%2 != 0 {
		count++
	}
	for i := low + 1; i < high; i++ {
		if i%2 == 1 {
			count++
		}
	}
	return count
}