package luoqiangLikou

import (
	"fmt"
	"testing"
)

func Test12(t *testing.T) {
	fmt.Println(intToRoman(20))
}

func intToRoman(num int) string {
	numArr := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	strArr := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	res := ""
	for num > 0 {
		for key, value := range numArr {
			if num >= value {
				res += strArr[key]
				num -= value
				break
			}
		}
	}

	return res
}
