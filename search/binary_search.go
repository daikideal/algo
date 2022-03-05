package search

func BinarySearchIterative(numbers []int, value int) (index int) {
	left, right := 0, len(numbers)-1

	for left < right {
		center := (left + right) / 2

		if numbers[center] == value {
			return center
		} else if numbers[center] < value {
			left = center + 1
		} else {
			right = center - 1
		}
	}

	return -1
}

func BinarySearchRecursive(numbers []int, value int) (index int) {
	var bsr func(numbers []int, value int, left int, right int) (index int)
	bsr = func(numbers []int, value int, left int, right int) (index int) {
		if left > right {
			return -1
		}

		center := (left + right) / 2
		if numbers[center] == value {
			return center
		} else if numbers[center] < value {
			return bsr(numbers, value, center+1, right)
		} else {
			return bsr(numbers, value, left, center+1)
		}
	}

	return bsr(numbers, value, 0, len(numbers)-1)
}
