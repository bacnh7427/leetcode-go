// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strings"
)

func main() {
	// t := []string{"g", "h", "i"}
	num := []int{1, 2, 0, 0}
	k := 34
	addToArrayForm(num, k)
}

func addToArrayForm(num []int, k int) []int {
	number := sliceToInt(num) + k
	fmt.Println(number)
	var split := strings.Split(number, "")
	fmt.Println(split)
	return split
}

func sliceToInt(s []int) int {
	res := 0
	op := 1
	for i := len(s) - 1; i >= 0; i-- {
		res += s[i] * op
		op *= 10
	}
	return res
}
