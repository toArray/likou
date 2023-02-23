package luoqiangLikou

import (
	"fmt"
	"sort"
	"testing"
)

func TestTopic611(T *testing.T) {
	//data := triangleNumber([]int{2, 2, 3, 4})
	//fmt.Println(data)

	arr := []int{4, 2, 3, 4}
	fmt.Println(triangleNumber(arr))
	//return

	//sort.Ints(arr)
	//count := 0
	//for k, v := range arr {
	//	for i := k + 1; k < len(arr); k++ {
	//		if sort.SearchInts(arr[k+1:], arr[i]+v) >= 0 {
	//			count++
	//		}
	//	}
	//}

}

func triangleNumber(nums []int) int {
	sort.Ints(nums)
	count := 0
	for i, _ := range nums {
		for j := i + 1; j < len(nums)-1; j++ {
			if search611(nums[j+1:], nums[j]) {
				count++
			}
		}
	}

	return count
}

func search611(arr []int, target int) bool {
	left, right := 0, len(arr)-1
	ans := false
	for left <= right {
		mid := int(uint(left+right) >> 1)
		if arr[mid] >= target {
			right = mid - 1
			ans = true
		} else {
			left = mid + 1
		}
	}

	return ans
}

//回溯超时
//func backtracking611(path []int, choose []int) {
//	//满足条件
//	if len(path) == 3 {
//		c := make([]int, 3)
//		copy(c, path)
//		if isValidTriangle(c) {
//			allPath = append(allPath, c)
//		}
//		return
//	}
//
//	for i := 0; i < len(choose); i++ {
//		path = append(path, choose[i])
//		backtracking611(path, choose[i+1:])
//		path = path[:len(path)-1]
//	}
//}
//
//func isValidTriangle(num []int) bool {
//	if num[0]+num[1] > num[2] && num[0]+num[2] > num[1] && num[1]+num[2] > num[0] {
//		return true
//	}
//	return false
//}
