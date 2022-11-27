package main

import "fmt"

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)
	ch3 := make(chan string, 1)

	ch1 <- "1"
	ch2 <- "1"
	ch3 <- "1"

	select {

	case message1 := <-ch1:
		fmt.Println("ch1 received:", message1)

	case message2 := <-ch2:
		fmt.Println("ch2 received:", message2)

	case message3 := <-ch3:
		fmt.Println("ch2 received:", message3)

	default:
		fmt.Println("No data received.")
	}
}
