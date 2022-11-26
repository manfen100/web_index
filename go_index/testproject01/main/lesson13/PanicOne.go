package main

import "fmt"

func panic(v interface{}) {}

func myTest() {
	defer fmt.Println("defer myTest")
	fmt.Println("we")
	panic("panic myTest")
}
func main() {
	defer fmt.Println("defer main")
	myTest()

	outOfArray(20)
}

func outOfArray(x int) {
	defer func() {
		//if err := recover(); err != nil {
		//	fmt.Println(err)
		//}
	}()

	var array [5]int
	array[x] = 1
}
