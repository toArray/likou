package luoqiangLikou

import (
	"fmt"
	"math"
	"testing"
)

func TestTopic633(T *testing.T) {
	fmt.Println(judgeSquareSum(5))
}

func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		res := left*left + right*right
		if res > c {
			right--
		} else if res < c {
			left++
		} else {
			return true
		}
	}

	return false
}
