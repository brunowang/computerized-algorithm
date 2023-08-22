package sorting

import (
	"math/rand"
)

func quickSortArray(nums []int) []int {
	shuffle(nums)
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	p := partition(nums, l, r)
	quickSort(nums, l, p-1)
	quickSort(nums, p+1, r)
}

func partition(nums []int, l, r int) int {
	p := nums[l]
	i, j := l+1, r
	if i > j {
		nums[l], nums[j] = nums[j], nums[l]
		return j
	}
	for {
		for i < r && nums[i] < p {
			i++
		}
		for j > l && nums[j] >= p {
			j--
		}
		if i >= j {
			break
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	nums[l], nums[j] = nums[j], nums[l]
	return j
}

func shuffle(nums []int) {
	n := len(nums)
	for i := 0; i < n; i++ {
		r := i + rand.Intn(n-i)
		nums[i], nums[r] = nums[r], nums[i]
	}
}
