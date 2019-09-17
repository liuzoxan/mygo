package template

// binary search template
func binarySearch(nums []int, t int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		// left + right big integer may overflow
		// mid := (left + right)/2
		mid := left + (right-left)/2
		if nums[mid] == t {
			// find the target!!

			return mid
		} else if nums[mid] > t {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return -1
}
