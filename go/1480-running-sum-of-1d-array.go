func runningSum(nums []int) []int {
	var result []int
	count := 0
	for i := 0; i < len(nums); i++ {
		count = nums[i] + count
		result = append(result, count)
	}
	return result
}