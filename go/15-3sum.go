func threeSum(nums []int) [][]int {
	result := [][]int{}
	if len(nums) < 3 {
		return result
	}

	sort.Ints(nums)

	for i, v := range nums {

		if i != 0 && nums[i-1] == v {
			continue
		}
		temp := twoSum(nums[i+1:], 0-v)
		if len(temp) > 0 {
			for i := range temp {
				result = append(result, append(temp[i], v))
			}
		}
	}
	return result
}

func twoSum(n []int, t int) [][]int {
	result := [][]int{}
	left := 0
	right := len(n) - 1
	for left < right {
		tt := n[left] + n[right]
		if tt == t {
			result = append(result, []int{n[left], n[right]})
			left++
			for left < len(n) && n[left] == n[left-1] {
				left++
			}
			right--

			for right >= 0 && n[right] == n[right+1] {
				right--
			}
			continue
		}
		if tt < t {
			left++
		} else {
			right--
		}
	}
	return result
}