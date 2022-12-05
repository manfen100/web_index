package main

import (
	"fmt"
	"strconv"
)

func cover() (a int) {
	x := 123
	y := fmt.Sprintf("%d", x)
	fmt.Println(y)
	z, h := fmt.Println(y, strconv.Itoa(x))
	fmt.Println(h)
	return z
}
func main() {
	x, _ := StrToInt("123")
	//y := strconv.Itoa(123)
	fmt.Println(x)
}

/*
字符串转数字
*/
func StrToInt(a string) (int, error) {
	// 字符串转数字
	x, err := strconv.Atoi("123")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(x)
	return x, err
}
