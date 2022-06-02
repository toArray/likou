package luoqiangLikou

import (
	"fmt"
	"strings"
	"testing"
)

func Test13(t *testing.T) {
	fmt.Println(RomanToInt("CMIIICM"))
}

func RomanToInt(s string) int32 {
	s = strings.ReplaceAll(s, "CM", "a")
	s = strings.ReplaceAll(s, "CD", "b")
	s = strings.ReplaceAll(s, "XC", "c")
	s = strings.ReplaceAll(s, "XL", "d")
	s = strings.ReplaceAll(s, "IX", "e")
	s = strings.ReplaceAll(s, "IV", "f")

	strMap := make(map[string]int32)
	strMap["M"] = 1000
	strMap["D"] = 500
	strMap["C"] = 100
	strMap["L"] = 50
	strMap["X"] = 10
	strMap["V"] = 5
	strMap["I"] = 1
	strMap["a"] = 900
	strMap["b"] = 400
	strMap["c"] = 90
	strMap["d"] = 40
	strMap["e"] = 9
	strMap["f"] = 4

	count := int32(0)
	for _, value := range s {
		count += strMap[string(value)]
	}

	return count
}
