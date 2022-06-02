package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test14(t *testing.T) {
	var strs = []string{"flower", "flow", "flight"}
	fmt.Println(longestCommonPrefix(strs))
}

func longestCommonPrefix(strs []string) string {
	index := 0
	str := ""

	for {
		//默认匹配
		tag := true

		//超出
		if index >= len(strs[0]) {
			break
		}

		//当前需要匹配的字段
		curStr := strs[0][index]
		for _, value := range strs {
			//当前字符串长度不够或者不匹配
			if index >= len(value) || value[index] != curStr {
				tag = false
				break
			}
		}

		if tag {
			//匹配则继续
			str += string(strs[0][index])
			index++
		} else {
			//不匹配直接退出
			break
		}
	}

	return str
}
