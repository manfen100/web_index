package main

import "fmt"

func slicearr() {
	var numList []int
	var numListEmpty = []int{}

	numList1 := make([]int, 3, 5)
	fmt.Println("", numList, numList1, numListEmpty)

	s := make([]string, 3, 5)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}
func mapmake() {
	scores := make(map[string]int)
	steps := make(map[string]string)

	var steps2 map[string]string = map[string]string{
		"1": "1",
	}

	steps3 := map[string]string{
		"": "",
	}

	steps3["第四部"] = "总监"

	delete(steps3, "第四部")

	v3, ok := steps3["第三部"]
	fmt.Println(ok)
	fmt.Println(v3)

	for key, value := range steps3 {
		fmt.Println("", key, value)
	}
	size := len(steps3)

	fmt.Println(scores, steps, steps2, steps3["第四部"], size)
}
