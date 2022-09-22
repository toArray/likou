package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic33(T *testing.T) {
	fmt.Println(search3([]int{4, 5, 6, 1, 2, 3}, 5))
}

func search3(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) >> 1
		if nums[mid] == target {
			return mid
		}

		if nums[0] <= nums[mid] {
			//左边有序
			if nums[0] <= target && target <= nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] <= target && target <= nums[len(nums)-1] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
