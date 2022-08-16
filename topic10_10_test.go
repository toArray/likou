package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test10_10(t *testing.T) {
	fmt.Println(isMatch("aa", ".*"))
}
func isMatch(s string, p string) bool {
	if p == ".*" {
		return true
	}

	if len(p) < 2 || string(p[len(p)-1]) != "*" {
		return p == s
	}

	first := p[:len(p)-1]
	for i := len(s) - 2; i >= 0; i-- {
		if string(s[i]) != first {
			return false
		}
	}

	return true
}
