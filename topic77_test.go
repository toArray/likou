package luoqiangLikou

import (
	"fmt"
	"testing"
)

//组合
func TestTopic77(T *testing.T) {
	backtracking(1, 5, 2)
	fmt.Println(res)
}

var (
	res  = make([][]int, 0)
	path = make([]int, 0)
)

//回溯
func backtracking(startIndex int, max int, l int) {
	if len(path) == l {
		c := make([]int, l)
		copy(c, path)
		res = append(res, c)
		return
	}

	for i := startIndex; i <= max; i++ {
		path = append(path, i)
		startIndex++
		backtracking(startIndex, max, l)
		path = path[:len(path)-1]
	}
}
