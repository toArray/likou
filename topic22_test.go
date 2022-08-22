package luoqiangLikou

import (
	"fmt"
	"testing"
)

var res22 = make([]string, 0)

//22. 括号生成
func Test22(t *testing.T) {
	fmt.Println(generateParenthesis(2))
}

func generateParenthesis(n int) []string {
	dfs("", n, 0, 0)
	return res22
}

func dfs(path string, n, left, right int) {
	if right > left || left > n {
		return
	}

	if len(path) == 2*n {
		res22 = append(res22, path)
		return
	}

	dfs(path+"(", n, left+1, right)
	dfs(path+")", n, left, right+1)
}
