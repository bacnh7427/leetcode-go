func numWaterBottles(numBottles int, numExchange int) int {
	var result int
	var emptyBottles int

	// first drink
	emptyBottles = numBottles
	result = numBottles

	for emptyBottles >= numExchange {
		// exchange
		numBottles = emptyBottles / numExchange
		emptyBottles = emptyBottles % numExchange

		// drink continue
		emptyBottles += numBottles
		result += numBottles
	}

	return result
}