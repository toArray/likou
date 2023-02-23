package luoqiangLikou

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestTopic658(T *testing.T) {

	fmt.Println(findClosestElements([]int{1, 2, 3, 4, 5}, 4, 3))
}

func findClosestElements(arr []int, k int, x int) []int {
	i := sort.SearchInts(arr, x)
	left := i - 1
	right := i

	count := 0
	var res []int
	for count < k {
		count++
		if left < 0 {
			res = append(res, arr[right])
			right++
			continue
		}
		if right == len(arr) {
			res = append(res, arr[left])
			left--
			continue
		}

		min := arr[left]
		r1 := math.Abs(float64(x - arr[right]))
		r2 := math.Abs(float64(x - arr[left]))
		if r1 < r2 {
			min = arr[right]
			right++
		} else {
			left--
		}
		res = append(res, min)
	}
	sort.Ints(res)
	return res
}
