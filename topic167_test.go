package luoqiangLikou

import "testing"

func TestTopic167(T *testing.T) {

}

func twoSum(numbers []int, target int) []int {
	for i := 0; i < len(numbers); i++ {
		left, right := i+1, len(numbers)-1
		for left <= right {
			mid := (left + right) >> 1
			if numbers[i]+numbers[mid] == target {
				return []int{i + 1, mid + 1}
			} else if numbers[mid] > (target - numbers[i]) {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}

	return nil
}
