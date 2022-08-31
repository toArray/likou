package luoqiangLikou

import (
	"fmt"
	"testing"
)

//x 的平方根
func TestTopic69(T *testing.T) {
	fmt.Println(mySqrt(1))
}

func mySqrt(x int) int {
	left, right := 0, x-1
	ans := -1
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
