package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test53(t *testing.T) {
	fmt.Println(maxSubArray([]int{1, 2, 3, 1}))
}

func maxSubArray(nums []int) int {
	sum := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] > sum {
			sum += nums[i]
		}
		if sum < nums[i] {
			sum = nums[i]
		}
	}

	return sum
}
