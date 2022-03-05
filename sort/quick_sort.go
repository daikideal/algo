package sort

// クイックソート
//
// ソート処理を実行する範囲を特定して分割関数を再帰的に呼び出す。
func QuickSort(numbers []int) []int {
	var _quickSort func(numbers []int, low int, high int)
	_quickSort = func(numbers []int, low int, high int) {
		if low < high {
			partition_index := partition(numbers, low, high)
			_quickSort(numbers, low, partition_index-1)
			_quickSort(numbers, partition_index+1, high)
		}
	}

	_quickSort(numbers, 0, len(numbers)-1)
	return numbers
}

// 分割関数
//
// ソート処理の実体。
// 任意の値pivotを境界値とし、pivot未満の値をリストの先頭側、
// pivot以上の値を終端側に分割する。
func partition(numbers []int, low int, high int) int {
	i := low - 1
	pivot := numbers[high]

	for j := low; j < high; j++ {
		if numbers[j] <= pivot {
			i++
			numbers[i], numbers[j] = numbers[j], numbers[i]
		}
	}
	numbers[i+1], numbers[high] = numbers[high], numbers[i+1]

	return (i + 1)
}
