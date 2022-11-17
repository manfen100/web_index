package main

import "fmt"

const (
	BEIJING  = 0
	SHANGHAI = 1
	SHENZHEN = 2
)

func main() {
	//fmt.Println("hello GoLang!")
	//
	//fmt.Println(len("你好"))                    //输出6，计算的是字节
	//fmt.Println(utf8.RuneCountInString("你好")) //输出2

	var a int
	fmt.Print("a=", a)
	fmt.Println("a的类型是：%T\n", a)

	var b int = 100
	fmt.Println("b =%d , type of b =%T\n", b)
	var bb string = "asa"
	fmt.Println("bb = %s, bb类型是 ：%T\n ", bb, bb)

	var c = 100
	var cc = "asasadad"
	fmt.Println("c,cc", c, cc)

	e := 100
	fmt.Println("e", e)
	f := "ew"
	fmt.Println("ew", f)

	var xx, yy int = 100, 200
	var kk, wx = 300, "21"
	var (
		nn int  = 100
		mm bool = true
	)
	fmt.Println("", xx, yy, kk, wx, nn, mm)

	//常量
	const length int = 10
	fmt.Println("length")

	fmt.Println("const=", BEIJING, SHANGHAI, SHENZHEN)

}
