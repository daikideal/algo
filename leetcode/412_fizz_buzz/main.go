package main

import (
	"strconv"
)

func fizzBuzz(n int) []string {
	answer := make([]string, n)

	for i := 1; i <= len(answer); i++ {
		var s string
		if i%15 == 0 {
			s = "FizzBuzz"
		} else if i%5 == 0 {
			s = "Buzz"
		} else if i%3 == 0 {
			s = "Fizz"
		} else {
			num := strconv.Itoa(i)
			s = num
		}
		answer[i-1] = s
	}

	return answer
}
