package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic1337(T *testing.T) {
	arr := [][]int{
		{1, 1, 1, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 1, 1, 0},
	}

	fmt.Println(kWeakestRows(arr, 1))
}

func kWeakestRows(mat [][]int, k int) []int {
	allCount := make([]int, len(mat), len(mat))
	for k, v := range mat {
		count := searchCount(v)
		allCount[k] = count
	}

	sortRes := make([]int, 0)
	count := 0
	for count <= len(mat[0]) {
		for i := 0; i < len(allCount); i++ {
			if allCount[i] == count {
				sortRes = append(sortRes, i)
			}
		}
		count++
	}

	return sortRes[:k]
}

func searchCount(arr []int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left + right) >> 1
		if arr[mid] >= 1 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return left
}
