package luoqiangLikou

import (
	"fmt"
	"testing"
)

//除自身以外数组的乘积
func TestTopic238(T *testing.T) {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}

func productExceptSelf(nums []int) []int {
	//记录左边的乘积
	res := make([]int, len(nums))
	left := 1
	for i := 0; i < len(nums); i++ {
		res[i] = left
		left = left * nums[i]
	}

	//记录右边的乘积
	right := 1
	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= right
		right *= nums[i]
	}

	return res
}
