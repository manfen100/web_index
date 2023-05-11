package main

import "fmt"

func main() {
	//
	//var x atomic.Int64
	//
	//go func() {
	//	x.Add(1) //增加
	//}()
	//
	//x.Store(x.Load() + 1) //load 读取、store存储

	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)
	}
}
