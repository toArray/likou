package luoqiangLikou

import (
	"fmt"
	"math"
	"testing"
)

func TestTopic07(T *testing.T) {
	res := topic07(-310)
	fmt.Println(res)
}

//整数反转
func topic07(x int) (res int) {
	for x != 0 {
		if x > math.MaxInt32 || x < -math.MaxInt32 {
			return 0
		}
		mod := x % 10
		x = x / 10
		res = res*10 + mod
	}

	return
}
