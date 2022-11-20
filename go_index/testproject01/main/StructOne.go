package main

import (
	"fmt"
)

//type struct_name struct {
//	attribute_name1 attribute_type
//	attribute_name2 attribute_type
//}

type Lesson struct {
	name   string
	target string
	spend  int
}

func main() {

	lesson1 := Lesson{
		name:   "",
		target: "",
		spend:  5,
	}

	lesson3 := struct {
		name, target string
		spend        int
	}{
		name:   "1",
		target: "1",
		spend:  3,
	}

	var lesson4 = Lesson{}

	fmt.Println(lesson1, lesson3, lesson4)

	lesson8 := &lesson3
	fmt.Println("dada", (*lesson8).name)
	fmt.Println("dada", lesson8.name)

	lesson10 := Lesson5{
		name:  "",
		spend: 1,
	}

	lesson10.author = Author{
		name: "",
		wx:   "",
	}
	fmt.Sprint(lesson10.name)
	fmt.Sprint(lesson10.author.wx)

	lesson11 := Lesson6{
		name:   "",
		target: "",
	}
	lesson11.Author = Author{
		name: "",
		wx:   "",
	}
	fmt.Sprint(lesson11.name)
	fmt.Sprint(lesson11.wx)
}

type Author struct {
	name string
	wx   string
}

type Lesson5 struct {
	name, target string
	spend        int
	author       Author
}

type Lesson6 struct {
	name, target string
	spend        int
	Author
}
