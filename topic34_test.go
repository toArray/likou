package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic34(T *testing.T) {
	//fmt.Println(searchRange([]int{5, 7, 7, 8, 8, 10}, 8))
	fmt.Println(searchRange([]int{1}, 1))
}

func searchRange(nums []int, target int) []int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			indexLeft, indexRight := mid, mid
			for nums[indexLeft-1] == target {
				indexLeft--
			}
			for nums[indexRight+1] == target {
				indexRight++
			}
			return []int{indexLeft, indexRight}
		}
	}

	return []int{-1, -1}
}
