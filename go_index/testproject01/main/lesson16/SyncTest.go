package lesson16

import "sync/atomic"

func main() {

	var x atomic.Int64

	go func() {
		x.Add(1) //增加
	}()

	x.Store(x.Load() + 1) //load 读取、store存储
}
