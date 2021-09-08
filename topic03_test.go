package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic03(T *testing.T) {
	s := "abcabcbb"
	s = "pwwkew"
	//s = "aaa"
	//s = "aaa"
	//s = "aaa"
	//s = "abbbb"
	res := topic03(s)
	fmt.Println(res)
}

//无重复字符的最长子串
func topic03(s string) int {
	windows := make(map[uint8]int) //滑动窗口
	left := 0                      //左指针
	right := 0                     //右指针
	res := 0                       //最大长度

	for i := 0; i < len(s); i++ {
		for {
			//是否在窗口内
			_, ok := windows[s[i]]
			if ok == false {
				break
			}

			//在窗口内左指针移动
			delete(windows, s[left])
			left++
		}

		//右指针移动
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
