package main

import (
	"fmt"
)

// TODO: DefinedTypeを使っていい感じにできないか？
type Stack struct {
	Stack []interface{}
}

func (s *Stack) Push(data interface{}) {
	s.Stack = append(s.Stack, data)

	return
}

func (s *Stack) Pop() interface{} {
	data := s.Stack[len(s.Stack)-1]
	s.Stack = s.Stack[:len(s.Stack)-1]

	return data
}

func main() {
	stack := new(Stack)
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	fmt.Println(stack)

	fmt.Println("==> poping...")
	fmt.Println(stack.Pop())
	fmt.Println(stack.Pop())
	fmt.Println(stack)
}
