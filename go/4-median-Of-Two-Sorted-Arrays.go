func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	aLength, bLength := len(nums1), len(nums2)
	cLength := aLength + bLength
	arr := []int{}

	i, j := 0, 0
	for i < aLength && j < bLength {
		a, b := nums1[i], nums2[j]
		if a <= b {
			arr = append(arr, a)
			i++
		}
		if a > b {
			arr = append(arr, b)
			j++
		}
	}

	for ; i < aLength; i++ {
		arr = append(arr, nums1[i])
	}
	for ; j < bLength; j++ {
		arr = append(arr, nums2[j])
	}

	if cLength%2 == 0 {
		return float64(arr[cLength/2-1]+arr[cLength/2]) / 2
	}

	return float64(arr[cLength/2])
}