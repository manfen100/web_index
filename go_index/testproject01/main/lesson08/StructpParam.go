package lesson08

import "fmt"

type Lesson struct {
	name, target string
	spend        int
}

func (l Lesson) ShowLessonInfo() {
	fmt.Println("name", l.name)
	fmt.Println("target", l.target)
	fmt.Println("spend", l.spend)
}

func (l *Lesson) AddTime(n int) {
	l.spend = l.spend + n
}

func main() {
	lesson13 := Lesson{
		name:   "1",
		target: "2",
		spend:  50,
	}
	fmt.Println("添加add方法前")
	lesson13.ShowLessonInfo()
	lesson13.AddTime(5)
	fmt.Println("添加add方法后")
	lesson13.ShowLessonInfo()
}
