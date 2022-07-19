package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test217(t *testing.T) {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 1}))
}

func containsDuplicate(nums []int) bool {
	hash := make(map[int]int32)
	for _, v := range nums {
		hash[v]++
		if hash[v] >= 2 {
			return true
		}
	}
	return false
}
