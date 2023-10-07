package lesson20

import "fmt"

func main() {
	var a int = 10
	var b int = 20
	swap(&a, &b)
	fmt.Println("a=", a, "b=", b)
}
func swap(pa *int, pb *int) {
	var temp int
	temp = *pa
	*pa = *pb
	*pb = temp

}
