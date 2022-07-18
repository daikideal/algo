package main

import (
	"strings"
	"testing"
)

func testReverseString(t *testing.T) {
	pattern := map[string]struct {
		inS  []string
		outS []string
	}{
		"A": {
			inS:  []string{"h", "e", "l", "l", "o"},
			outS: []string{"o", "l", "l", "e", "h"},
		},
		"B": {
			inS:  []string{"H", "a", "n", "n", "a", "h"},
			outS: []string{"h", "a", "n", "n", "a", "H"},
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T) {
			input := strings.Join(v.inS, "")

			actual := []byte(input)
			reverseString([]byte(actual))

			output := []byte(strings.Join(v.outS, ""))
			for i := range actual {
				if actual[i] != output[i] {
					t.Error()
				}
			}
		})
	}
}
