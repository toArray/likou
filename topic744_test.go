package luoqiangLikou

import "testing"

func TestTopic744(T *testing.T) {

}

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right := 0, len(letters)-1
	ans := letters[0]
	for left <= right {
		mid := (left + right) / 2
		if letters[mid] >= target {
			ans = letters[mid]
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return ans
}
