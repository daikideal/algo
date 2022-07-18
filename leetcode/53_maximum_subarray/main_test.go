package main

import (
	"testing"
)

func testMaxSubArray(t *testing.T) {
	pattern := map[string]struct {
		inNums []int
		outNum int
	}{
		"A": {
			inNums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			outNum: 6,
		},
		"B": {
			inNums: []int{1},
			outNum: 1,
		},
		"C": {
			inNums: []int{5, 4, -1, 7, 8},
			outNum: 23,
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T) {
			actual := maxSubArray(v.inNums)

			if actual != v.outNum {
				t.Error()
			}
		})
	}
}
