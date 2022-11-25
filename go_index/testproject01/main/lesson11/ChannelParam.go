package main

import "fmt"

//// 通道
//var ch chan string
//
////ch = make(chan string)
//
//ab := make(chan string)
//
//ab <- data
//
//value := <- ab

func PrintChan(c chan string) {
	// 往通道传入数据
	c <- "123"
}
func main() {
	//创建通道
	ch := make(chan string)
	fmt.Println("dayin")

	// 开启写协程
	go PrintChan(ch)
	rec := <-ch

	fmt.Println(rec)

	close(ch)

	c := make(chan int, 3)
	//传入
	c <- 1
	c <- 2
	//取出
	<-c

}
