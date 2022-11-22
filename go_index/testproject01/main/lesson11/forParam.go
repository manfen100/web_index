package main

import "fmt"

func fortest() {
	num := 0
	for num < 4 {
		fmt.Println(num)
		num++
	}

	for num := 0; num < 4; num++ {
		fmt.Println(num)
	}

	str := "123"
	for index, value := range str {
		fmt.Println(index, value)
	}

	for {

	}
	for {

	}
}
func bookPrint() {
	fmt.Println("go")
}
func main() {
	defer bookPrint()

	fmt.Println("")
}
