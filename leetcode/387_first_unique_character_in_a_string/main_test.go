package main

import (
	"testing"
)

func testfirstUniqChar(t *testing.T) {
	pattern := map[string]struct {
		inString string
		outIndex int
	}{
		"A": {
			inString: "leetcode",
			outIndex: 0,
		},
		"B": {
			inString: "loveleetcode",
			outIndex: 2,
		},
		"C": {
			inString: "aabb",
			outIndex: -1,
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T) {
			actual := firstUniqChar(v.inString)
			if actual != v.outIndex {
				t.Errorf("Expect is %d, actual is %d", v.outIndex, actual)
			}
		})
	}
}
