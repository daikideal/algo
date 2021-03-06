package sort

func BubbleSort(numbers []int) []int {
	len_numbers := len(numbers)

	for i := 0; i < len_numbers; i++ {
		for j := 0; j < (len_numbers - i - 1); j++ {
			if numbers[j] > numbers[j+1] {
				numbers[j], numbers[j+1] = numbers[j+1], numbers[j]
			}
		}
	}

	return numbers
}
