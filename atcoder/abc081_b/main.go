package main

import (
	"fmt"
)

func main() {
	var length int
	fmt.Scanf("%d", &length)

	nums := make([]int, length)
	for i, _ := range nums {
		fmt.Scanf("%d", &nums[i])
	}

	var count int
	var flag bool
	for {
		flag = false

		for _, n := range nums {
			if n%2 == 0 {
				flag = true
			} else {
				flag = false
				break
			}
		}

		if !flag {
			break
		}

		count++
		for i, _ := range nums {
			nums[i] = nums[i] / 2
		}
	}

	fmt.Print(count)
}
