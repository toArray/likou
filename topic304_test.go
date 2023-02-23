package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic304(T *testing.T) {
	data := make([][]int, 0)
	data = append(data, []int{5, 6, 3, 2, 1})
	data = append(data, []int{1, 2, 0, 1, 5})
	data = append(data, []int{4, 1, 0, 1, 7})
	data = append(data, []int{1, 0, 3, 0, 5})
	r := Constructor(data)
	r.SumRegion(2, 1, 4, 3)
	fmt.Println(r)
}

type NumMatrix struct {
	matrix [][]int
}

func Constructor(matrix [][]int) NumMatrix {
	dp := make([][]int, len(matrix))
	for k, v := range matrix {
		dp[k] = make([]int, len(v)+1)
		for i := 0; i < len(v); i++ {
			dp[k][i+1] = dp[k][i] + v[i]
		}
	}

	return NumMatrix{
		matrix: dp,
	}
}

func (this *NumMatrix) SumRegion(row1 int, col1 int, row2 int, col2 int) int {
	sum := 0
	for i := row1; i <= row2; i++ {
		sum += this.matrix[i][col2+1] - this.matrix[i][col1]
	}
	return sum
}
