package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic153(T *testing.T) {
	fmt.Println(findMin([]int{1, 2, 3}))
}

func findMin(nums []int) int {
	left, right := 0, len(nums)-1
	for left < right {
		mid := (left + right) >> 1
		if nums[mid] >= nums[right] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return nums[left]
}
