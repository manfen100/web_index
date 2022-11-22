package main

import "fmt"

type Book struct {
	bookName, authorName string
}

func (b Book) printName() {
	fmt.Println("--", b.bookName, b.authorName)
}

func main() {
	//book := Book{"书", "hh"}
	//defer book.printName()
	//fmt.Println("main...")
	//
	//defer fmt.Printf("1")
	//defer fmt.Printf("2")
	//defer fmt.Printf("3")
	//fmt.Printf("main...----------------")

	lesson := showLesson()
	fmt.Println("main: s =", s)
	fmt.Println("main: lesson =", lesson)

}

var s string = "123"

func showLesson() string {
	defer func() {
		s = "456"
	}()
	fmt.Println("showLesson:s= ", s)
	return s
}
