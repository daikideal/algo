package sort

import (
	"math/rand"
	"time"
)

// 引数で受け取った長さのランダムなintのスライスを返す
func RandomNumbers(length int) []int {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, length)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	return numbers
}
