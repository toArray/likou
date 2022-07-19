package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test3471(t *testing.T) {
	var strs = []int{54, 42, 62, 75, 82, 86, 86}
	fmt.Println(findMaxCI(strs))
}

func findMaxCI(nums []int) int {
	l := len(nums)
	if l == 0 {
		return 0
	}

	max := 0
	cur := 0
	for i := 0; i < l; i++ {
		if i+1 != l {
			if nums[i+1] > nums[i] {
				cur++
				if cur > max {
					max = cur
				}
			} else {
				cur = 0
			}
		}
	}

	max += 1
	return max
}
