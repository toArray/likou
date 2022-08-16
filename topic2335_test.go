package luoqiangLikou

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

type Str string

func Test2335(t *testing.T) {
	var test Str
	test = "这是一个字符串"

	bb, _ := json.Marshal(test) //byte
	fmt.Println(string(bb))

	//fmt.Println(fillCups([]int{1, 2, 0}))
}

//144

//装满杯子需要的最短总时长
func fillCups(amount []int) int {
	count := 0
	fmt.Println(1 / 2)
	sort.Ints(amount)
	if amount[0]+amount[1] > amount[2] {
		count = (amount[0] + amount[1] + amount[2] + 1) / 2
	} else {
		count = amount[2]
	}

	return count
}
