package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic74(T *testing.T) {
	//arr := [][]int{
	//	{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60},
	//}
	arr := [][]int{
		{1, 3},
	}
	fmt.Println(searchMatrix(arr, 3))
}

func searchMatrix(matrix [][]int, target int) bool {
	left, right := 0, len(matrix)-1
	var res []int
	for left <= right {
		mid := (left + right) >> 1
		temp := matrix[mid]
		first := temp[0]
		last := temp[len(temp)-1]

		if target >= first && target <= last {
			res = temp
			break
		} else if target > last {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	l, r := 0, len(res)-1
	for l <= r {
		mid := (l + r) >> 1
		if target < res[mid] {
			r = mid - 1
		} else if target > res[mid] {
			l = mid + 1
		} else {
			return true
		}
	}

	return false
}
