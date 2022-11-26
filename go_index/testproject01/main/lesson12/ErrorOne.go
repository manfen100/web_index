package lesson12

import (
	"fmt"
)

func main() {
	//a := 100
	//b := -10
	//r, err := area(a, b)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println("Area=", r)

	a := 100
	b := -10
	area, err := rectangleArea(a, b)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("length %d or width %d is less than zero", err.length, err.width)
			return
		}
		fmt.Println(err)
		return
	}
	fmt.Println("Area =", area)
}

func area(a, b int) (int, error) {
	if a < 0 || b < 0 {
		//return 0, errors.New("错了")
		return 0, fmt.Errorf("", a, b)
	}
	return a * b, nil
}

type areaError struct {
	err    string
	length int
	width  int
}

func (e *areaError) Error() string {
	return fmt.Sprintln("", e.err, e.length, e.width)
}

func rectangleArea(a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, &areaError{"", a, b}
	}
	return a * b, nil
}
