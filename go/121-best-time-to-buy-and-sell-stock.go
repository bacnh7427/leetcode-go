func maxProfit(prices []int) int {
	min := math.MaxInt32
	result := 0
	for _, v := range prices {
		if v < min {
			min = v
		} else if v-min > result {
			result = v - min
		}
	}
	return result
}