package luoqiangLikou

import "testing"

func TestTopic441(T *testing.T) {

}

func arrangeCoins(n int) int {
	//前N等差求和公式 x * (x + 1) / 2

	left, right := 0, n
	ans := 0
	for left <= right {
		mid := (left + right) >> 1   //多少行
		need := mid * (mid + 1) >> 1 //需要多少个
		if need > n {
			right = mid - 1
		} else {
			ans = mid
			left = mid + 1
		}
	}
	return ans
}
