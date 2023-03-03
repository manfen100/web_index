package main

import "fmt"

//import "github.com/gin-gonic/gin"
//
//func main() {
//	r := gin.Default()
//	r.GET("/ping", func(context *gin.Context) {
//		context.JSON(200, gin.H{"message": "pong"})
//	})
//	r.Run()
//}

func main() {
	func() {
		for i := 0; i < 300; i++ {
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
	//	for j := 0; j < 300; j++ {
	//		defer func() {
	//			fmt.Println("b:", j)
	//		}()
	//	}
	//}()
}
