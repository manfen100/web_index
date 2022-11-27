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
}
