func minCostClimbingStairs(cost []int) int {
	for i := 2; i < len(cost); i++ {
		if cost[i-1] < cost[i-2] {
			cost[i] += cost[i-1]
		} else {
			cost[i] += cost[i-2]
		}
	}
	result := cost[len(cost)-1]
	if result > cost[len(cost)-2] {
		result = cost[len(cost)-2]
	}
	return result
}