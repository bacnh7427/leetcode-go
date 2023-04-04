func findJudge(n int, trust [][]int) int {
	if len(trust) == 0 && n == 1 {
		return 1
	}
	result := -1
	countPeopleTrustedYou := make(map[int]int)
	countYouTrustedPeople := make(map[int]int)
	for _, t := range trust {
		countPeopleTrustedYou[t[1]]++
		countYouTrustedPeople[t[0]]++
	}

	// countPeopleTrustedYou: key: 1 nguoi trong lang, value: so nguoi ma tin tuong vao nguoi nay
	for key, value := range countPeopleTrustedYou {
		if value == n-1 && countYouTrustedPeople[key] == 0 {
			result = key
			break
		}
	}
	return result
}