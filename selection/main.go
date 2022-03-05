package main

import (
	"fmt"
	"math/rand"
)

func SelectionSort(numbers []int) []int {
	len_numbers := len(numbers)

	for i := 0; i < len_numbers; i++ {
		min_index := i

		for j := (i + 1); j < len_numbers; j++ {
			if numbers[min_index] > numbers[j] {
				numbers[min_index], numbers[j] = numbers[j], numbers[min_index]
			}
		}
	}

	return numbers
}

func main() {
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println(SelectionSort(numbers))
}
