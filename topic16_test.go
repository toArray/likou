package luoqiangLikou

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

//最接近三数之和
func Test16(t *testing.T) {
	nums := []int{1, 1, 1, 0}
	fmt.Println(threeSumClosest(nums, 100))
}

//双指针
func threeSumClosest(nums []int, target int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	fmt.Println(nums)
	temp := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums); i++ {
		var left, right = i + 1, len(nums) - 1
		for left < right {
			threeSum := nums[left] + nums[right] + nums[i]
			if int(math.Abs(float64(threeSum-target))) < int(math.Abs(float64(temp-target))) {
				temp = threeSum
			}
			if threeSum > target {
				right--
			} else {
				left++
			}
		}
	}

	return temp
}
