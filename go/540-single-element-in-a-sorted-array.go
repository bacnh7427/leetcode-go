func singleNonDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		j := i + 1
		valueLeft := nums[i]
		if j == len(nums) {

			return nums[j-1]

		} else {
			valueRight := nums[j]
			if valueLeft != valueRight {
				return nums[i]
			}
			i++
		}

	}
	return nums[0]
}