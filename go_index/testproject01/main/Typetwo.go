package main

import "fmt"

// 使用切片
func changeSlice(value []int) {
	value[0] = 200
}

// 使用数组指针
func changeArray(value *[3]int) {
	(*value)[0] = 200
}

func main() {
	x := [3]int{1, 2, 3}
	changeSlice(x[:])
	fmt.Println(x)

	y := [3]int{1, 2, 3}
	changeArray(&y)

}
