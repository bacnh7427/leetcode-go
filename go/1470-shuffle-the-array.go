func shuffle(nums []int, n int) []int {
	result := make([]int, len(nums))
	for i := 0; i < n; i++ {
		result[i*2] = nums[i]
		result[i*2+1] = nums[i+n]
	}
	return result
}