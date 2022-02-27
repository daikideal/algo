package main

import "fmt"

type BinarySearchTree struct {
	Root *Node
}

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func (tree *BinarySearchTree) Insert(value int) {
	if tree.Root == nil {
		tree.Root = &Node{Value: value}
		return
	}

	var insert func(node *Node, value int) *Node
	insert = func(node *Node, value int) *Node {
		if node == nil {
			newNode := &Node{Value: value}
			return newNode
		}

		if node.Value > value {
			node.Left = insert(node.Left, value)
		} else {
			node.Right = insert(node.Right, value)
		}

		return node
	}

	insert(tree.Root, value)
	return
}

func (tree *BinarySearchTree) Inorder() {
	var inorder func(node *Node)
	inorder = func(node *Node) {
		if node != nil {
			inorder(node.Left)
			fmt.Println(node.Value)
			inorder(node.Right)
		}

		return
	}

	inorder(tree.Root)
	return
}

func (tree *BinarySearchTree) Search(value int) bool {
	if tree.Root == nil {
		return false
	}

	var search func(node *Node, value int) bool
	search = func(node *Node, value int) bool {
		if node == nil {
			return false
		}

		if node.Value == value {
			return true
		} else if node.Value > value {
			return search(node.Left, value)
		} else if node.Value < value {
			return search(node.Right, value)
		}

		return false
	}

	result := search(tree.Root, value)
	return result
}

func (tree *BinarySearchTree) Remove(value int) *Node {
	if tree.Root == nil {
		return nil
	}

	var remove func(node *Node, value int) *Node
	remove = func(node *Node, value int) *Node {
		if node == nil {
			return node
		}

		if node.Value > value {
			node.Left = remove(node.Left, value)
		} else if node.Value < value {
			node.Right = remove(node.Right, value)
		} else {
			if node.Left == nil {
				return node.Right
			} else if node.Right == nil {
				return node.Left
			}

			minNode := tree.minValue(node.Right)
			node.Value = minNode.Value
			node.Right = remove(node.Right, value)
		}

		return node
	}

	replacedNode := remove(tree.Root, value)
	return replacedNode
}

// 引数で渡したNodeのうち、最小の値を持つNodeへのポインタを返す
func (tree *BinarySearchTree) minValue(node *Node) *Node {
	currentNode := node
	for currentNode.Left != nil {
		currentNode = currentNode.Left
	}

	return currentNode
}

func main() {
	tree := new(BinarySearchTree)
	tree.Insert(5)
	tree.Insert(3)
	tree.Insert(2)
	tree.Insert(10)
	tree.Insert(7)

	tree.Inorder()
	fmt.Println(tree.Search(2))
	fmt.Println(tree.Search(4))

	fmt.Println("==> Removing...")
	tree.Remove(2)
	fmt.Println(tree.Search(2))
}
