package main

import (
	"fmt"
	"time"
)

type Sender = chan<- string

type Receiver = <-chan string

func main() {
	var ch = make(chan string)

	go func() {
		var sender Sender = ch
		fmt.Println("1")
		sender <- "gp"
	}()

	go func() {
		var reveiver Receiver = ch
		message := <-reveiver
		fmt.Println("dd", message)

	}()
	time.Sleep(time.Millisecond)

	var ch2 = make(chan int, 5)
	go loopPrint(ch2)

	for v := range ch2 {
		fmt.Println(v)
	}

	ch3 := make(chan bool, 1)

	var x int

	for i := 0; i < 10000; i++ {
		go increment(ch3, &x)
		time.Sleep(time.Millisecond)
		fmt.Println("x =", x)
	}
}

func loopPrint(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	// 记得要关闭通道
	// 否则主协程遍历完不会结束，而会阻塞
	close(c)
}

func increment(ch chan bool, x *int) {
	ch <- true
	*x = *x + 1
	<-ch
}
