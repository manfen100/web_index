package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//死锁
	ch := make(chan bool)
	go funcReceive(ch)
	ch <- true
	time.Sleep(time.Millisecond)

	// 实例化 sync.WaitGroup
	var waitGroup sync.WaitGroup
	//传入子携程数量
	waitGroup.Add(3)
	go task(1, &waitGroup)
	go task(2, &waitGroup)
	go task(3, &waitGroup)
	waitGroup.Wait()
}

func funcReceive(c chan bool) {
	fmt.Println(<-c)
}

func task(taskNum int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Println(taskNum, i)
	}
}
