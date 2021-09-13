package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic05(T *testing.T) {
	s := "amadam"
	//s = "ada"
	//s = "addad"
	s = "add"
	s = "aadd"
	res := topic05(s)
	fmt.Println(res)
}

//最长回文子串中心扩散法
func topic05(s string) string {
	if len(s) <= 1 {
		return s
	}

	res := string(s[0])
	l := len(s)

	for i := 0; i < l; i++ {
		res = odd(s, i, res)
		res = even(s, i, res)
	}

	return res
}

func even(s string, i int, res string) string {
	right := i + 1 //右边位置
	left := i      //左边位置定义为自己

	for {
		//超出边界不继续比对
		if left < 0 || right >= len(s) {
			break
		}

		//左右相等字符更新最长回文
		if s[left] == s[right] {
			backStr := s[left : right+1]
			if len(backStr) > len(res) {
				res = backStr
			}

			left--
			right++
		} else {
			break
		}
	}

	return res
}

func odd(s string, i int, res string) string {
	right := i + 1 //右边位置
	left := i - 1  //左边位置

	for {
		//超出边界不继续比对
		if left < 0 || right >= len(s) {
			break
		}

		//左右相等字符更新最长回文
		if s[left] == s[right] {
			backStr := s[left : right+1]
			if len(backStr) > len(res) {
				res = backStr
			}

			left--
			right++
		} else {
			break
		}
	}

	return res
}
