package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic04(T *testing.T) {
	s := "abcabcbb"
	s = "pwwkew"
	//s = "aaa"
	//s = "aaa"
	//s = "aaa"
	//s = "abbbb"
	res := topic04(s)
	fmt.Println(res)
}

//无重复字符的最长子串
func topic04(s string) int {
	windows := make(map[uint8]int)
	left := 0
	right := 0
	res := 0

	for i := 0; i < len(s); i++ {
		for {
			_, ok := windows[s[i]]
			if ok == false {
				break
			}

			delete(windows, s[left])
			left++
		}

		res = getMax(res, right-left+1)
		right++
		windows[s[i]]++
	}

	return res
}

func getMax(a, b int) int {
	max := a
	if b > a {
		max = b
	}
	return max
}
