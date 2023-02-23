package luoqiangLikou

import (
	"fmt"
	"math"
	"testing"
)

func TestTopic209(T *testing.T) {
	fmt.Println(minSubArrayLen(7, []int{2, 3, 1, 2, 4, 3}))
}

func minSubArrayLen2(target int, nums []int) int {
	left, right := 0, 0
	min := math.MaxInt32
	sum := 0

	for right < len(nums) {
		sum += nums[right]
		for sum >= target {
			if right-left+1 < min {
				min = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}

	if min == math.MaxInt32 {
		return 0
	}

	return min
}

func minSubArrayLen(target int, nums []int) int {
	left, right := 0, 0
	min := 0
	windows := []int{}
	sum := 0

	for left <= right {
		if sum >= target {
			sum -= nums[left]
			windows = windows[1:]
			left++
		} else {
			if right == len(nums) {
				break
			}
			sum += nums[right]
			windows = append(windows, nums[right])
			right++
		}

		if len(windows) <= min && sum >= target || min == 0 && sum >= target {
			min = len(windows)
			if min == 1 {
				break
			}
		}
	}

	return min
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
