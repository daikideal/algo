package main

import (
	"testing"
)

func testFizzBuzz(t *testing.T) {
	pattern := map[string]struct {
		inNum      int
		outStrings []string
	}{
		"A": {
			inNum:      3,
			outStrings: []string{"1", "2", "Fizz"},
		},
		"B": {
			inNum:      5,
			outStrings: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		"C": {
			inNum:      15,
			outStrings: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T) {
			actual := fizzBuzz(v.inNum)

			for i := range actual {
				if actual[i] != v.outStrings[i] {
					t.Error()
				}
			}
		})
	}
}
