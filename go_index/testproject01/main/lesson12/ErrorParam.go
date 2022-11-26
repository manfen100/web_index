package lesson12

import (
	"fmt"
	"os"
)

type error interface {
	Error() string
}

func main() {
	file, err := os.Open("/a.txt")

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(file.Name(), "")
}

type errString struct {
	s string
}

func (e *errString) Error() string {
	return e.s
}
func New(text string) error {
	return &errString{text}
}
