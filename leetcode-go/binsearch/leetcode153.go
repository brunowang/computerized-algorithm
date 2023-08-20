package binsearch

func findMin(nums []int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l < r {
		mid := l + (r-l)/2
		if nums[r] < nums[mid] {
			l = mid + 1
		} else if nums[mid] < nums[r] {
			r = mid
		} else if nums[mid] == nums[r] {
			r--
		}
	}
	return nums[r]
}
