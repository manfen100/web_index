package lesson21

import "fmt"

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(fact(7))
	var fib func(n int) int
	fib = func(n int) int {
		if n < 2 {
			return n
		}
		return fib(n - 1)
	}

	//指针
	i := 1
	fmt.Println("initial:", i)
	zeroval(i)

	zeroptr(&i)

	//结构体
	newPerson("name")
	s := person{name: "Sean", age: 50}
	sp := &s
	sp.name = "Sean"

}
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}

}

func fact(n int) int {
	if n == 0 {
		return 1
	}
	return n * fact(n-1)
}

func zeroval(ival int) {
	ival = 0
}
func zeroptr(iptr *int) {
	*iptr = 0
}
func examineRune(r rune) {
	if r == 't' {
		fmt.Printf("")
	} else if r == 'n' {
		fmt.Printf("")
	}
}

type person struct {
	name string
	age  int
}

func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}
