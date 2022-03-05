package main

import (
	"math/rand"
	"testing"
	"time"
)

func RandomNumbers(length int) []int {
	rand.Seed(time.Now().UnixNano())

	numbers := make([]int, length)
	for i := range numbers {
		numbers[i] = rand.Intn(100)
	}

	return numbers
}

func TestMergeSort(t *testing.T) {
	list := RandomNumbers(10)
	t.Logf("origin: %v", list)

	actual := MergeSort(list)
	for i := 0; i < len(actual)-1; i++ {
		if actual[i] > actual[i+1] {
			t.Errorf("actual:%v", actual)
		}
	}

	t.Logf("actual: %v", actual)
}
