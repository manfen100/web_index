package main

import (
	"fmt"
)

func PrintInfo() {
	fmt.Println("从0到Go语言微服务架构师")
}

func main() {
	// 开启一个协程执行 PrintInfo 函数
	go PrintInfo()
	// 使主协程休眠 1 秒
	///time.Sleep(1 * time.Second)
	// 打印 main
	fmt.Println("main")
}
