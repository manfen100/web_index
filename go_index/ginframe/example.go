package main

import (
	"log"
	"os"
)

func main() {
	//
	//r := gin.Default()
	//r.GET("/ping", func(c *gin.Context) {
	//	c.JSON(200, gin.H{
	//		"message": "pong",
	//	})
	//})
	//r.Run()

	tom := Person{
		Name:   "zhangsan",
		Age:    18,
		Logger: log.New(os.Stdout, "", 0),
	}

	tom.Printf("hello %s", "world")
}

type Person struct {
	Name string
	Age  int
	*log.Logger
}

func (p *Person) Printf(format string, v ...any) {
	print("hello")

}
