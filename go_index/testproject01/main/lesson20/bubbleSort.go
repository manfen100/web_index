package lesson20

func bubbleSort(numbers []int) {
	n := len(numbers)

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}
}

func main() {
	numbers := []int{5, 4, 3, 2, 1}
	bubbleSort(numbers)
	for _, v := range numbers {
		println(v)
	}
}
