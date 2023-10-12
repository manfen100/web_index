package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	server := NewServer("127.0.0.1", 8888)
	server.Start()
}
