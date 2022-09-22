package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic1885(T *testing.T) {
	var a = []int{55, 30, 5, 4, 2}
	var b = []int{100, 20, 10, 10, 5}
	fmt.Println(maxDistance(a, b))
	fmt.Println(search4(b, 1111))
}

func maxDistance(nums1 []int, nums2 []int) int {
	max := 0
	for i := 0; i < len(nums1); i++ {
		ans := search4(nums2, nums1[i])
		if ans == -1 {
			continue
		}
		if ans-i > max {
			max = ans - i
		}
	}
	return max
}

func search4(nums1 []int, target int) int {
	left, right := 0, len(nums1)-1
	ans := -1
	for left <= right {
		mid := (left + right) >> 1
		if nums1[mid] >= target {
			ans = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return ans
}
