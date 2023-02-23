package luoqiangLikou

import (
	"fmt"
	"testing"
)

func TestTopic1812(T *testing.T) {
	str := "c7"
	res := !((str[0]%2 == 1 && str[1]%2 == 1) || (str[0]%2 == 0 && str[1]%2 == 0))
	fmt.Println(res)

}
