package luoqiangLikou

import (
	"encoding/json"
	"fmt"
	"sort"
	"testing"
)

//三数之和
func Test15(t *testing.T) {
	nums := []int{3, 0, -2, -1, 1, 2}
	fmt.Println(threeSumDoublePoint(nums))
}

//双指针
func threeSumDoublePoint(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	data := make([][]int, 0, 0)
	m := make(map[string]struct{})
	sort.Ints(nums)
	fmt.Println(nums)
	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			break
		}

		var tag = -nums[i]
		var left, right = i + 1, len(nums) - 1
		for left < right {
			twoSum := nums[left] + nums[right]
			if tag == twoSum {
				r := []int{nums[i], nums[left], nums[right]}
				sort.Ints(r)
				b, _ := json.Marshal(r)
				if _, ok := m[string(b)]; ok == false {
					data = append(data, r)
					m[string(b)] = struct{}{}
				}
				left++
				right--
			} else if twoSum < tag {
				left++
			} else {
				right--
			}
		}
	}

	return data
}

//暴力
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}

	data := make([][]int, 0, 0)
	m := make(map[string]struct{})
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == 0 {
					r := []int{nums[i], nums[j], nums[k]}
					sort.Ints(r)
					b, _ := json.Marshal(r)
					if _, ok := m[string(b)]; ok == false {
						data = append(data, r)
						m[string(b)] = struct{}{}
					}
				}
			}
		}
	}

	return data
}
