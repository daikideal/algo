package doubly_linked_list

import (
	"fmt"
)

type DoublyLinkedList struct {
	Head *Node
}

type Node struct {
	Data interface{}
	Prev *Node
	Next *Node
}

func (list *DoublyLinkedList) Append(data interface{}) {
	newNode := &Node{Data: data}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	currentNode := list.Head
	for currentNode.Next != nil {
		currentNode = currentNode.Next
	}

	currentNode.Next = newNode
	newNode.Prev = currentNode

	return
}

func (list *DoublyLinkedList) Insert(data interface{}) {
	newNode := &Node{Data: data}

	if list.Head == nil {
		list.Head = newNode
		return
	}

	list.Head.Prev = newNode
	newNode.Next = list.Head
	list.Head = newNode

	return
}

func (list *DoublyLinkedList) Remove(data interface{}) {
	currentNode := list.Head

	if currentNode != nil && currentNode.Data == data {
		list.Head = currentNode.Next
		currentNode = nil
		return
	}

	for currentNode != nil && currentNode.Data != data {
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return
	}

	if currentNode.Next == nil {
		previousNode := currentNode.Prev
		previousNode.Next = nil
		currentNode = nil
		return
	} else {
		previousNode := currentNode.Prev
		nextNode := currentNode.Next
		previousNode.Next = nextNode
		nextNode.Prev = previousNode
		currentNode = nil
		return
	}
}

func (list *DoublyLinkedList) Reverse() {
	previousNode := new(Node)
	currentNode := list.Head

	for currentNode != nil {
		previousNode = currentNode.Prev
		currentNode.Prev = currentNode.Next
		currentNode.Next = previousNode

		currentNode = currentNode.Prev
	}

	if previousNode != nil {
		list.Head = previousNode.Prev
	}

	return
}

func (list DoublyLinkedList) Print() {
	currentNode := list.Head

	for currentNode != nil {
		fmt.Println(currentNode.Data)
		currentNode = currentNode.Next
	}

	return
}
