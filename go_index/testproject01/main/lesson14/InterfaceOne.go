package main

import "fmt"

func SHowType(i interface{}) {
	fmt.Println("", i, i)
}
func main() {
	str := ""
	SHowType(str)
	num := 3.14
	SHowType(num)

	x := "hello"
	for _, x := range x {
		x := x + 'A' - 'a'
		fmt.Printf("%c", x) // "HELLO" (one letter per iteration)
	}
}

var a string
