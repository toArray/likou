package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic35(T *testing.T) {
	fmt.Println(searchInsert([]int{1, 3, 3, 3, 3, 5, 6}, 3))
}

func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	ans := -1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] >= target {
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		}
		ans = right
	}

	return ans + 1
}
