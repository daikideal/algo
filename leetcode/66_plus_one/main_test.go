package main

import (
	"testing"
)

func testPlusOne(t *testing.T) {
	pattern := map[string]struct {
		inDigits   []int
		outNums  []int
	}{
		"When nums is sorted": {
			inDigits:   []int{1,2,3},
			outNums:  []int{1, 2, 4},
		},
		"When nums is not sorted": {
			inDigits:   []int{4, 3, 2, 1},
			outNums:  []int{4,3, 2, 2},
		},
		"When nums is same": {
			inDigits:   []int{9},
			outNums:  []int{1, 0},
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T){
			actual := plusOne(v.inDigits)

			for i := range actual {
				if actual[i] != v.outNums[i] {
					t.Error()
				}
			}
		})
	}
}
