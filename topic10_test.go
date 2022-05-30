package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test10(t *testing.T) {
	arr := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(arr))
}

//盛最多的水
func maxArea(height []int) int {
	count := len(height) - 1
	left := 0
	right := count
	max := 0

	maxLeftIndex := 0
	maxRightIndex := 0

	for left != right {
		a := height[left]
		b := height[right]
		min := a
		if b < a {
			min = b
		}

		temp := min * count
		if temp > max {
			max = temp
			maxLeftIndex = left
			maxRightIndex = right
		}

		count--
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	fmt.Println(maxLeftIndex, maxRightIndex)
	return max
}
