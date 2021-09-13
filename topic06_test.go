package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic06(T *testing.T) {
	s := "AB"
	res := topic06(s, 1)
	fmt.Println(res)
}

//Z字形变换
func topic06(s string, numRows int) string {
	if len(s) <= 1 || numRows == 1 {
		return s
	}

	base := numRows*2 - 2
	data := make([][]byte, numRows)
	for i := 0; i < len(s); i++ {
		mod := i % base
		if mod < numRows {
			data[mod] = append(data[mod], s[i])
		} else {
			data[base-mod] = append(data[base-mod], s[i])
		}
	}

	res := []byte{}
	for _, value := range data {
		res = append(res, value...)
	}

	return string(res)
}
