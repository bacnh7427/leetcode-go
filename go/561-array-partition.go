func arrayPairSum(nums []int) int {
	result := 0
	sort.Ints(nums)
	for i := 0; i < len(nums); i = i + 2 {
		result += nums[i]
	}
	return result
}