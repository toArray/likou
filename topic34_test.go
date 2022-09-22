package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic34(T *testing.T) {
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1, 2, 2, 3}, 2))
}

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			if nums[mid] == target {
				ans = mid
			}
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	ans2 := -1
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) >> 1
		if nums[mid] > target {
			r = mid - 1
		} else {
			if nums[mid] == target {
				ans2 = mid
			}
			l = mid + 1
		}
	}

	return []int{ans, ans2}
}
