package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic1346(T *testing.T) {
	//fmt.Println(checkIfExist([]int{-20, 8, -6, -14, 0, -19, 14, 4}))
	fmt.Println(checkIfExist([]int{-2, 0, 10, -19, 4, 6, -8}))
}

func checkIfExist(arr []int) bool {
	//每个数x2倍
	res := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		res[arr[i]*2]++
	}

	for i := 0; i < len(arr); i++ {
		c, ok := res[arr[i]]
		if arr[i] == 0 {
			if c > 1 {
				return true
			}
		} else {
			if ok {
				return ok
			}
		}
	}

	return false
}
