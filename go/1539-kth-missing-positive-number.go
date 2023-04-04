func findKthPositive(arr []int, k int) int {
	l := 0
	lenArr := len(arr)
	for l < lenArr {
		mid := l + (lenArr-l)/2
		if arr[mid]-(mid+1) >= k {
			lenArr = mid
		} else {
			l = mid + 1
		}
	}
	return l + k
}