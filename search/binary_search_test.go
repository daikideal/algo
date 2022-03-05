package search

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	pattern := map[string]struct {
		inList   []int
		inValue  int
		outValue int
	}{
		"When the value is found, should return the value": {
			inList:   []int{0, 1, 2, 3, 4, 5},
			inValue:  2,
			outValue: 2,
		},
		"When the value is not found, should return `-1`": {
			inList:   []int{0, 1, 2, 3, 4, 5},
			inValue:  10,
			outValue: -1,
		},
		// TODO: `inList`がソートされていない場合のテストはどう書く？
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T) {
			actual := BinarySearchIterative(v.inList, v.inValue)
			if actual != v.outValue {
				t.Errorf("Failed with Iterative. Expect: %v", v.outValue)
			}

			actual = BinarySearchRecursive(v.inList, v.inValue)
			if actual != v.outValue {
				t.Errorf("Failed with Recursive. Expect: %v", v.outValue)
			}
		})
	}
}
