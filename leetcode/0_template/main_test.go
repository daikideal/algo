package main

import (
	"testing"
)

func testTemplate(t *testing.T) {
	pattern := map[string]struct{
		input interface{}
		output interface{}
	}{
		"A": {
			input: 0,
			output: 0,
		},
	}

	for k, v := range pattern {
		t.Run(k, func(t *testing.T){
			// write here test
			t.Errorf("Expect is %v", v)
		})
	}
}
