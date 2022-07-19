package luoqiangLikou

import (
	"fmt"
	"testing"
)

//反转字符串中的单词 III
func TestTopic557(T *testing.T) {
	s := "Let's take LeetCode contest"
	res := reverseWords(s)
	fmt.Println(res)
}

func reverseWords(s string) string {
	tempArr := make([]int32, 0)
	res := ""
	for k, value := range s {
		if string(value) == " " {
			for i := len(tempArr) - 1; i >= 0; i-- {
				res += string(tempArr[i])
			}
			if string(value) == " " {
				res += " "
			}
			tempArr = []int32{}
		} else {
			tempArr = append(tempArr, value)
			if k == len(s)-1 {
				for i := len(tempArr) - 1; i >= 0; i-- {
					res += string(tempArr[i])
				}
			}
		}
	}
	return res
}
