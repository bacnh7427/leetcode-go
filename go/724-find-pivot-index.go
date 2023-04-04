func pivotIndex(nums []int) int {
	sumNums := 0
	for _, value := range nums {
		sumNums += value
	}

	sumLeft := 0
	for i, value := range nums {
		if sumLeft == sumNums-sumLeft-value {
			return i
		}
		sumLeft += value
	}
	return -1
}