package main

import (
	"fmt"
	"math/rand"
)

// マージソート
//
// リストの長さが2以上の場合、リストを2つに分割する。
// 分割したリストをマージ処理へと渡す。
func mergeSort(numbers []int) []int {
	if len(numbers) < 2 {
		return numbers
	}

	center := len(numbers) / 2
	left := numbers[:center]
	right := numbers[center:]

	left, right = mergeSort(left), mergeSort(right)

	return merge(left, right)
}

// マージ関数
//
// 2つのリストの先頭同士を比較し、小さい方を結果のリストに格納する。
// これを2つのリストが両方空になるまで再帰的に処理する。
func merge(left []int, right []int) []int {
	result := []int{}
	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	for ; i < len(left); i++ {
		result = append(result, left[i])
	}

	for ; j < len(right); j++ {
		result = append(result, right[j])
	}

	return result
}

func main() {
	numbers := make([]int, 10)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	fmt.Println(mergeSort(numbers))
}
