package list

import (
	"fmt"
)

type LinkedList struct {
	Head *Node
}

type Node struct {
	Data interface{}
	Next *Node // 初期値をnilとする
}

// リストの終端にデータを追加する
func (list *LinkedList) Append(data interface{}) {
	newNode := &Node{Data: data}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	lastNode := list.Head
	for lastNode.Next != nil {
		lastNode = lastNode.Next
	}
	lastNode.Next = newNode
	return
}

// リストの先端にデータを追加する
func (list *LinkedList) Insert(data interface{}) {
	newNode := &Node{Data: data}
	newNode.Next = list.Head

	list.Head = newNode
	return
}

// 引数と最初に一致したデータをリストから削除する
func (list *LinkedList) Remove(data interface{}) {
	currentNode := list.Head

	if currentNode != nil && currentNode.Data == data {
		list.Head = currentNode.Next
		currentNode = nil
		return
	}

	previousNode := new(Node)
	for currentNode != nil && currentNode.Data != data {
		previousNode = currentNode
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return
	}

	previousNode.Next = currentNode.Next
	currentNode = nil
	return
}

func (list *LinkedList) Reverse() {
	previousNode := new(Node)
	currentNode := list.Head

	for currentNode != nil {
		nextNode := currentNode.Next
		currentNode.Next = previousNode
		// リストの終端が全フィールドがnilのNodeになることを防ぐ
		if currentNode.Next.Data == nil {
			currentNode.Next = nil
		}
		previousNode = currentNode
		currentNode = nextNode
	}
	list.Head = previousNode

	return
}

// リスト内のデータを全て出力する
func (list LinkedList) Print() {
	currentNode := list.Head

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	return
}
