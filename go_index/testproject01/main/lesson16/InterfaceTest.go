package main

import (
	"fmt"
)

type Programmer interface {
	WriteHelloWord() string
}

type GoProgrammer struct {
	name  string
	price float64
}

func (g *GoProgrammer) WriteHelloWord() string {
	fmt.Println("implent")

	return g.name
}
func main() {
	//var p Programmer
	//p = new(GoProgrammer)
	//
	//

	p := GoProgrammer{
		name:  "1",
		price: 1.0,
	}
	s := p.WriteHelloWord()
	fmt.Println("s->:", s)
	fmt.Println(s)
}
