// Sử dụng tìm kiếm nhị phân
// Chia đôi mảng
// So sánh số cần tìm kiếm với số có vị trí ở giữa xem số đó nằm bên trái hay bên phải.
// Sử dụng vòng lặp và so sánh tiếp...

func searchInsert(nums []int, target int) int {
	leftNum := 0
	rightNum := len(nums) - 1
	for leftNum <= rightNum {
		mid := leftNum + (rightNum-leftNum)/2
		if nums[mid] < target {
			leftNum = mid + 1
		} else if nums[mid] > target {
			rightNum = mid - 1
		} else {
			return mid
		}

	}
	return leftNum
}