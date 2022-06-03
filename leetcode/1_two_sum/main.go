package main

import (
	"fmt"
)

// GOOD: Using Hash
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int, len(nums))

	for i, num := range nums {
		// num と足し合わせてtargetになる値がmemoに存在するか確認し、
		// あればそれが解として返る、なければmemoに追加する
		//
		// 参考: https://stackoverflow.com/questions/2050391/how-to-check-if-a-map-contains-a-key-in-go
		if index, ok := memo[target-num]; ok {
			return []int{index, i}
		}
		memo[num] = i
	}

	return []int{0, 0}
}

// BAD: bruteforce
// func twoSum(nums []int, target int) []int {
// 	ans := make([]int, 2)

// 	for i := range nums {
// 		for j := 0; j < len(nums)-1; j++ {
// 			if i == j {
// 				continue
// 			}

// 			if (nums[i] + nums[j]) == target {
// 				ans = []int{i, j}
// 				break
// 			}
// 		}
// 	}

// 	return ans
// }

func main() {
	pattern := []struct {
		nums   []int
		target int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9},
		{nums: []int{3, 2, 4}, target: 6},
		{nums: []int{3, 3}, target: 6},
	}

	for _, v := range pattern {
		fmt.Printf("input: %v, output: %v", v, twoSum(v.nums, v.target))
		fmt.Println()
	}
}
