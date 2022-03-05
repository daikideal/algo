package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	list := RandomNumbers(10)
	t.Logf("origin: %v", list)

	actual := SelectionSort(list)
	for i := 0; i < len(actual)-1; i++ {
		if actual[i] > actual[i+1] {
			t.Errorf("actual:%v", actual)
		}
	}

	t.Logf("actual: %v", actual)
}
