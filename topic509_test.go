package luoqiangLikou

import (
	"fmt"
	"testing"
)

//斐波那契
func Test509(t *testing.T) {
	fmt.Println(fib(1))
}

func fib(n int) int {
	if n <= 0 {
		return 0
	}

	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	if n < 2 {
		return 1
	}

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
