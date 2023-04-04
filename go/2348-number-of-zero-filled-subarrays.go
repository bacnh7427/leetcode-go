func zeroFilledSubarray(nums []int) int64 {
	var count, total uint
	for _, num := range nums {
		if num != 0 {
			count = 0
		} else {
			count += 1
		}

		total += count
	}
	return int64(total)
}