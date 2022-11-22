package main

import "fmt"

type Lesson struct {
	Name      string
	Target    string
	SpendTime int
}

func (lesson Lesson) PrintInfo() {
	fmt.Println("name:", lesson.Name)
	fmt.Println("target:", lesson.Target)
	fmt.Println("spendTime:", lesson.SpendTime)
}

func (lesson Lesson) changeLessonName(name string) {
	lesson.Name = name
}

func (lesson *Lesson) AddSpendTime(n int) {
	lesson.SpendTime = lesson.SpendTime + n
}
