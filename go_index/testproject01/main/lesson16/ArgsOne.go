package main

import (
	"fmt"
	"strconv"
	"time"
)

type IntConv func(op int) int

func timeSpent(inner IntConv) IntConv {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)

		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func TestGrouptine() {
	for i := 0; i < 10; i++ {
		go func(i int) { fmt.Println(i) }(i)
	}

}

func slowFunc(op int) int {
	time.Sleep(time.Second * 1)
	fmt.Println("oper")
	return op
}

func TestConv() {
	s := strconv.Itoa(10)
	fmt.Println(s)
	if i, err := strconv.Atoi("10"); err == nil {
		fmt.Println(i)
	}

}

func Sum(ops ...int) int {
	ret := 0
	for _, op := range ops {
		ret += op
	}
	return ret
}

func Clear() {
	fmt.Println("clear")
}

func Defer() {
	defer Clear()
	fmt.Println("start")
}

func main() {
	//if len(os.Args) > 1 {
	//	fmt.Println("hello", os.Args[1])
	//}
	tssF := timeSpent(slowFunc)
	fmt.Println(tssF(10))

	Sum(1, 2, 3)
	Sum(1)

}
