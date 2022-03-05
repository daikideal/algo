package sort

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
