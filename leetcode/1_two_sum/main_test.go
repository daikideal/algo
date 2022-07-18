package main

import (
	"testing"
)

func testTwoSum(t *testing.T) {
	pattern := map[string]struct {
		inNums   []int
		inTarget int
		outNums  []int
	}{
		"When nums is sorted": {
			inNums:   []int{2, 7, 11, 15},
			inTarget: 9,
			outNums:  []int{0, 1},
		},
		"When nums is not sorted": {
			inNums:   []int{3, 2, 4},
			inTarget: 6,
			outNums:  []int{1, 2},
		},
		"When nums is same": {
			inNums:   []int{3, 3},
			inTarget: 6,
			outNums:  []int{0, 1},
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T) {
			actual := twoSum(v.inNums, v.inTarget)

			for i := range actual {
				if actual[i] != v.outNums[i] {
					t.Error()
				}
			}
		})
	}
}
