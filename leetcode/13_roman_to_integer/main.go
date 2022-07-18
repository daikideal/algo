package main

func romanToInt(s string) int {
	m := map[int32]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var sum int
	var previous int32

	for i := len([]rune(s)) - 1; i >= 0; i-- {
		if num, ok := m[[]rune(s)[i]]; ok {
			if num < m[previous] {
				sum = sum - num
			} else {
				sum = sum + num
			}

			previous = []rune(s)[i]
		}
	}

	return sum
}
