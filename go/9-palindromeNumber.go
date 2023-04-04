func isPalindrome(x int) bool {
	iReverse := 0
	for iNum := x; iNum > 0; iNum /= 10 {
		iReverse = iReverse*10 + (iNum % 10)
	}

	return iReverse == x
}