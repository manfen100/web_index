package main

import "fmt"

func showChar() {
	var x byte = 65
	var y uint8 = 65
	fmt.Println("x", x, y)
}

func sizeOfChar() {
	var x byte = 65
	fmt.Println("x", x)
}

//	func main() {
//		showChar()
//		sizeOfChar()
//
//		var study string
//		study = "1"
//		study2 := "3"
//	}
func test01() {
	var arr [5]int
	fmt.Println(arr)

	arr[0] = 1
}

func test02() {
	var arr1 = [5]int{1, 2, 3, 4, 5}

	arr2 := [5]int{1, 2, 3, 4, 5}

	arr3 := [5]int{1, 2}

	arr4 := [5]int{1: 100, 4: 200}

	arr5 := [...]int{1, 2, 3}

	fmt.Println(arr1, arr2, arr3, arr4, arr5)
}

func arrLength() {
	arr := [...]string{"1"}
	fmt.Println("", arr, len(arr))
}

func showArr() {
	arr := [...]string{"1"}
	for index, value := range arr {
		fmt.Println("", index, value)
	}
}
func showArr1() {
	arr := [...]string{"1"}
	for _, value := range arr {
		fmt.Println("", value)
	}
}
