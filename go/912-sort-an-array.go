func sortArray(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}

	mid := len(nums) / 2

	sortedLeft := sortArray(nums[:mid])
	sortedRight := sortArray(nums[mid:])

	return merge(sortedLeft, sortedRight)
}

func merge(a, b []int) []int {
	merged := make([]int, 0)

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			merged = append(merged, a[i])
			i++
		} else {
			merged = append(merged, b[j])
			j++
		}
	}

	if i < len(a) {
		merged = append(merged, a[i:]...)
	}

	if j < len(b) {
		merged = append(merged, b[j:]...)
	}

	return merged
}