package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)

	go func() {
		fmt.Println("send :1")
		c <- 1
	}()

	// 以并发的方式调用匿名函数
	go func() {
		n := <-c
		fmt.Println("receive:", n)
	}()

	time.Sleep(time.Millisecond)
}
func chanReceve() {
	//定义只读通道
	//c:=make(chan string)
	////
	//type Receiver = <-chan string
	//var receiver Receiver = c

	type Receiver = <-chan int
	receiver := make(Receiver)
	fmt.Println(receiver)

	//c:=make(chan int)
	//type Sender = chan <-int
	//var sender Sender = c

	type Sender = chan<- int
	sender := make(Sender)
	fmt.Println(sender)
}
