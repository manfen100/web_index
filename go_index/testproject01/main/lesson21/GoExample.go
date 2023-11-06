package lesson21

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		println(i)
		i = i + 1
	}
	for j := 7; j <= 9; j++ {
		println(j)
	}
	for {
		print()
	}
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		println(n)
	}

	var a [5]int
	a[4] = 100
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println(b)

	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	m := make(map[string]int)
	m["k1"] = 7
	delete(m, "k1")

	_, prs := m["k2"]
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("prs", prs)
	fmt.Println("map", n)

	nums := []int{2, 3, 4}
	sum := 0
	for _, num := range nums {
		sum += num
	}

	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf(k, v)
	}

	sum1(1, 2)
	numa := []int{1, 2, 3, 4}
	sum1(numa...)

}
func plus(a int, b int) int {
	return a + b
}
func plusPlus(a, b, c int) int {
	return a + b + c
}
func vals() (int, int) {
	return 3, 7
}
func sum1(nums ...int) {
	fmt.Print(nums, "")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
