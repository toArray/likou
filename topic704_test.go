package luoqiangLikou

import (
	"fmt"
	"testing"
)

//二分查找
func TestTopic704(T *testing.T) {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 12
	fmt.Println(search2(nums, target))
}

/*
search
@Param 二分法-左闭右闭
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := (left + right) / 2
		num := nums[mid]
		if num > target {
			right = mid - 1
		} else if num < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}

/*
search
@Param 二分法-左闭右开
*/
func search2(nums []int, target int) int {
	left, right := 0, len(nums)

	for left < right {
		mid := (left + right) / 2
		num := nums[mid]
		if num > target {
			right = mid
		} else if num < target {
			left = mid + 1
		} else {
			return mid
		}
	}
	return -1
}
