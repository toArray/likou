package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic06(T *testing.T) {
	s := "PAYPALISHIRING"
	res := topic06(s, 3)
	fmt.Println(res)
}

//Z字形变换
func topic06(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	cur := 0
	flag := false
	data := make([][]byte, numRows)
	for i := 0; i < len(s); i++ {
		data[cur] = append(data[cur], s[i])
		if cur == 0 || cur == numRows-1 {
			flag = !flag
		}

		if flag {
			cur++
		} else {
			cur--
		}
	}

	res := []byte{}
	for _, value := range data {
		res = append(res, value...)
	}

	return string(res)
}
