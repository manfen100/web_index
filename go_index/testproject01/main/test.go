package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	fmt.Println("hello GoLang!")

	fmt.Println(len("你好"))                    //输出6，计算的是字节
	fmt.Println(utf8.RuneCountInString("你好")) //输出2
}
