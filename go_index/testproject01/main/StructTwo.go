package main

import "fmt"

type Lessson121 struct {
	name, targrt string
	spend        int
}

func (l Lessson121) ShowLessonInfo() {
	fmt.Println("name", l.name)
	fmt.Println("tagret", l.targrt)
}
func main() {

}
