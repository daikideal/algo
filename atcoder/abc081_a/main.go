package main

import (
	"fmt"
)

func main() {
	var in string

	fmt.Scanf("%s", &in)

	var count int
	for _, c := range in {
		if string(c) == "1" {
			count++
		}
	}

	fmt.Print(count)
}
