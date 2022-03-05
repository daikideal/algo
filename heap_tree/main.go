package main

import (
	"fmt"
	"math"
)

type MiniHeapTree struct {
	Heap        []int
	CurrentSize int
}

// システムで一番小さい数字
//
// ヒープにデータを入れていく基になる数字。
const sysMin int = -1 * math.MaxInt

func NewMiniHeap() *MiniHeapTree {
	heap := make([]int, 0)
	heap = append(heap, sysMin)

	return &MiniHeapTree{
		Heap:        heap,
		CurrentSize: 0,
	}
}

func (tree *MiniHeapTree) Push(value int) {
	tree.Heap = append(tree.Heap, value)
	tree.CurrentSize += 1
	tree.heapifyUp(tree.CurrentSize)

	return
}

func (tree *MiniHeapTree) Pop() int {
	if len(tree.Heap) == 1 {
		return -1
	}

	root := tree.Heap[1]
	last := tree.Heap[len(tree.Heap)-1]
	tree.Heap = tree.Heap[:len(tree.Heap)-1]
	if len(tree.Heap) == 1 {
		// treeを再構成する必要がないためこの時点でreturn
		return root
	}

	tree.Heap[1] = last
	tree.CurrentSize -= 1
	tree.heapifyDown(1)

	return root
}

func (tree *MiniHeapTree) Print() {
	fmt.Println(tree.Heap[1:])
}

func (MiniHeapTree) parentIndex(index int) int {
	return index / 2
}

func (MiniHeapTree) leftchildIndex(index int) int {
	return index * 2
}

func (MiniHeapTree) rightchildIndex(index int) int {
	return (index * 2) + 1
}

func (tree *MiniHeapTree) swap(index1 int, index2 int) {
	tree.Heap[index1], tree.Heap[index2] = tree.Heap[index2], tree.Heap[index1]
	return
}

// 指定したノードを上へ追いやる
func (tree *MiniHeapTree) heapifyUp(index int) {
	// 指定したノードが先頭でない限り
	for tree.parentIndex(index) > 0 {
		if tree.Heap[index] < tree.Heap[tree.parentIndex(index)] {
			tree.swap(index, tree.parentIndex(index))
		}

		index = tree.parentIndex(index)
	}

	return
}

// 指定したノードを下へ追いやる
func (tree *MiniHeapTree) heapifyDown(index int) {
	// 次のノードが存在する限り
	for tree.leftchildIndex(index) <= tree.CurrentSize {
		minChildIndex := tree.minChildIndex(index)
		if tree.Heap[index] > tree.Heap[minChildIndex] {
			tree.swap(index, minChildIndex)
		}

		index = minChildIndex
	}

	return
}

func (tree MiniHeapTree) minChildIndex(index int) int {
	if tree.rightchildIndex(index) > tree.CurrentSize {
		return tree.leftchildIndex(index)
	}

	if tree.Heap[tree.leftchildIndex(index)] < tree.Heap[tree.rightchildIndex(index)] {
		return tree.leftchildIndex(index)
	} else {
		return tree.rightchildIndex(index)
	}
}

func main() {
	heap := NewMiniHeap()
	heap.Push(3)
	heap.Push(8)
	heap.Push(1)
	heap.Push(2)
	heap.Push(5)
	heap.Push(10)
	heap.Print()

	fmt.Println("==> Poping...")
	fmt.Println(heap.Pop())
	heap.Print()
	fmt.Println(heap.Pop())
	heap.Print()
}
