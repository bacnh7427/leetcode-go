func canPlaceFlowers(flowerbed []int, n int) bool {
	lenFlo := len(flowerbed)
	for i := 0; i < lenFlo && n > 0; i += 2 {
		if flowerbed[i] == 0 {
			if i+1 == lenFlo || flowerbed[i+1] == 0 {
				n--
			} else {
				i++
			}
		}
	}
	if n == 0 {
		return true
	}
	return false
}