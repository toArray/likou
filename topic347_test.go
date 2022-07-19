package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test347(t *testing.T) {
	var strs = []int{1, 1, 2, 3, 4, 5, 5}
	fmt.Println(topKFrequent(strs, 1))
}

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int64)
	res := make([]int, 0)
	for _, value := range nums {
		count[value]++
		if len(res) == 0 {
			res = append(res, value)
		}

	}

	return res
}
