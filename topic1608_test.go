package luoqiangLikou

import (
	"fmt"
	"sort"
	"testing"
)

func TestTopic1608(T *testing.T) {
	fmt.Println(specialArray([]int{3, 5}))
}

func specialArray(nums []int) int {
	sort.Ints(nums)
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		count := len(nums[mid:])
		if count == nums[mid] {
			return mid
		} else if nums[mid] > count {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
