package main

import "fmt"

func tyoe1() {
	x := "q"
	ptr := &x

	ptr2 := new(string)
	*ptr2 = "123"

	fmt.Println("----", ptr2)
	fmt.Println("----", *ptr2)

	fmt.Println("www", ptr)

	x2 := "声明变量"
	var p *string
	p = &x2

	fmt.Println(p)

	x1 := "从0到Go语言微服务架构师"
	var ptr1 *string
	ptr1 = &x1
	fmt.Println("ptr is ", *ptr1)
}

func changeByPointer(value *int) {
	*value = 200
}

func main() {
	tyoe1()
	x3 := 99
	p3 := &x3
	fmt.Println("执行changeByPointer函数之前p3是", *p3)

	changeByPointer(p3)
	fmt.Println("执行changeByPointer函数之后p3是", *p3)
}
