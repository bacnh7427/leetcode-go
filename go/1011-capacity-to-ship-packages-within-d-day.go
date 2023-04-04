func shipWithinDays(weights []int, days int) int {
	var low int = 0
	var high int = 1000000000

	for low < high {
		mid := low + (high-low)/2
		if ok(weights, mid, days) {
			high = mid
		} else {
			low = mid + 1
		}
	}

	return low
}

func ok(weights []int, capacity int, days int) bool {
	var cnt int = 1
	var currentSum int = 0
	for _, weight := range weights {
		if weight > capacity {
			return false
		}
		if currentSum+weight > capacity {
			cnt++
			currentSum = 0
		}
		currentSum += weight
	}
	return cnt <= days
}
