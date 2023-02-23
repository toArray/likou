package luoqiangLikou

import (
	"fmt"
	"strings"
	"testing"
)

func TestTopic459(T *testing.T) {

	fmt.Println(repeatedSubstringPattern("abab"))
}

func repeatedSubstringPattern(s string) bool {
	if len(s)%2 == 1 {
		return false
	}

	str := ""
	for len(s) != 0 {
		if strings.HasPrefix(s, str) {
			if len(s) == len(str) {
				return true
			}
		} else {
			str += string(s[0])
		}
		s = s[1:]
	}

	return false
}
