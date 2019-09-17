package template


// binary search template
func binarySearch(nums []int, t int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right)/2
		if nums[mid] == t {
			// find the target!!

			return mid
		} else if nums[mid] > t {
			right = mid-1
		} else {
			left = mid +1
		}
	}

	return -1
}