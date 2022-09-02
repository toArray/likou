package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic1539(T *testing.T) {
	fmt.Println(findKthPositive([]int{2, 3}, 2))
	//fmt.Println(findKthPositive([]int{1, 2, 3, 4}, 2))
}

func findKthPositive(arr []int, k int) int {
	res := 1
	index := 0
	for k > 0 && index < len(arr) {
		if res < arr[index] {
			k--
		} else {
			index++
		}
		res++
	}
	return res + k - 1
}
