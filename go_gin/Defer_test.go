package main

import "fmt"

func main() {
	func() {
		for i := 0; i < 3; i++ {
			defer fmt.Println("a:", i)
		}
	}()

	func() {
		for i := 0; i < 300; i++ {
			defer fmt.Println("b:", i)
		}
	}()
	//fmt.Println()
	//func() {
	//	for i := 0; i < 3; i++ {
	//		defer func() {
	//			fmt.Println("b:", i)
	//		}()
	//	}
	//}()

}
