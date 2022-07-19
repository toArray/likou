package luoqiangLikou

import (
	"fmt"
	"testing"
	"time"
)

func Test15(t *testing.T) {
	threeSum()
}

func threeSum() {
	fmt.Println(1)

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	go func() {
		fmt.Println(2)
		panic("3")
	}()
	time.Sleep(time.Second)
	fmt.Println(4)
}
