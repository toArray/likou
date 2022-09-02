package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic1531(T *testing.T) {
	arr := [][]int{
		{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3},
	}
	fmt.Println(countNegatives(arr))
}

func countNegatives(grid [][]int) int {
	count := 0
	for _, v := range grid {
		left, right := 0, len(v)-1
		for left <= right {
			mid := (left + right) >> 1
			if v[mid] >= 0 {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

		count += len(v) - left
	}
	return count
}
