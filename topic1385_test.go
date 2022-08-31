package luoqiangLikou

import (
	"fmt"
	"math"
	"sort"
	"testing"
)

func TestTopic1385(T *testing.T) {
	arr1 := []int{4, 5, 8}
	arr2 := []int{10, 9, 1, 8}
	d := 2
	res := findTheDistanceValue(arr1, arr2, d)
	fmt.Println(res)
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	sort.Ints(arr1)
	sort.Ints(arr2)
	count := 0

	for j := 0; j < len(arr1); j++ {
		flag := true
		for i := 0; i < len(arr2); i++ {
			abs := int(math.Abs(float64(arr1[j] - arr2[i])))
			if abs <= d {
				flag = false
				break
			}
		}
		if flag {
			count++
		}
	}

	return count
}
