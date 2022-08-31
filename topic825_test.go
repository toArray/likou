package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic825(T *testing.T) {
	res := peakIndexInMountainArray([]int{0, 1, 0})
	fmt.Println(res)
}

func peakIndexInMountainArray(arr []int) int {
	left, right := 0, len(arr)
	for left < right {
		mid := (left + right) / 2
		//山峰刚好最中间
		if arr[mid] > arr[mid-1] && arr[mid] > arr[mid+1] {
			return mid
		}

		if arr[mid+1] > arr[mid] {
			left = mid //右边比左边高，说明山峰在右侧
		} else if arr[mid+1] < arr[mid] {
			right = mid //右边比左边低，山峰在左侧
		}

		//if arr[mid] > arr[mid+1] {
		//	right = mid //左边山峰比较大，这边不-1，要保留最大山峰
		//} else if arr[mid] < arr[mid+1] {
		//	left = mid + 1 //左边山峰比较小，移动左山峰
		//} else if arr[mid] > arr[mid-1] {
		//	left = mid
		//} else if arr[mid] < arr[mid-1] {
		//	right = mid - 1
		//} else {
		//	return mid
		//}

	}
	return -1
}
