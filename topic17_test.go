package luoqiangLikou

import (
	"fmt"
	"testing"
)

//电话号码的字母组合
func Test17(t *testing.T) {
	fmt.Println(letterCombinations("23"))
}

var (
	res17  []string
	path17 string
	words  = make(map[string][]string)
)

func init() {
	words["2"] = []string{"a", "b", "c"}
	words["3"] = []string{"d", "e", "f"}
	words["4"] = []string{"g", "h", "i"}
	words["5"] = []string{"j", "k", "l"}
	words["6"] = []string{"m", "n", "o"}
	words["7"] = []string{"p", "q", "r", "s"}
	words["8"] = []string{"t", "u", "v"}
	words["9"] = []string{"w", "x", "y", "z"}
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return res17
	}
	back(0, path17, digits)
	return res17
}

func back(start int, path string, digits string) {
	if len(path) == len(digits) {
		res17 = append(res17, path)
		return
	}

	key := digits[uint8(start)]
	content := words[string(key)]
	start++
	for j := 0; j < len(content); j++ {
		path += content[j]
		back(start, path, digits)
		path = path[:len(path)-1]
	}
}
