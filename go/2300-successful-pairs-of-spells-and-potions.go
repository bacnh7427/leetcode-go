// using binary search
func successfulPairs(spells []int, potions []int, success int64) []int {
	sort.Ints(potions)
	lspells := len(spells)
	lpotions := len(potions)
	ans := make([]int, lspells)
	for i := 0; i < lspells; i++ {
		left := 0
		right := lpotions

		for left < right {
			mid := left + (right-left)/2
			if int64(potions[mid]*spells[i]) >= success {
				right = mid
			} else {
				left = mid + 1
			}
		}
		ans[i] = lpotions - left
	}
	return ans
}