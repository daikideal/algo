package list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	pattern := map[string]struct {
		Head   *Node
		inData interface{}
	}{
		"When head is nil, should set value to head": {
			Head:   nil,
			inData: 1,
		},
		"When head is not nil, should set value to last": {
			Head:   &Node{Data: 1, Next: nil},
			inData: 2,
		},
	}

	for k, v := range pattern {
		list := new(LinkedList)
		if v.Head != nil {
			list.Head = v.Head
		}

		t.Run(k, func(t *testing.T) {
			list.Append(v.inData)

			actual := list.Head
			for actual.Next != nil {
				actual = actual.Next
			}

			if actual.Data != v.inData {
				t.Errorf("Expected is {%v}, but actual is {%v}", v.inData, actual)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	pattern := map[string]struct {
		Head   *Node
		inData interface{}
	}{
		"Should set value to head": {
			Head:   nil,
			inData: 1,
		},
	}

	for k, v := range pattern {
		list := new(LinkedList)

		t.Run(k, func(t *testing.T) {
			list.Insert(v.inData)

			actual := list.Head.Data
			if actual != v.inData {
				t.Errorf("Expected is {%v}, but actual is %v", v.inData, actual)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	pattern := map[string]struct {
		NodeValues []interface{}
		inData     interface{}
	}{
		"When passed value is in list": {
			NodeValues: []interface{}{1, 2, 3},
			inData:     1,
		},
		"When passed value is not in list": {
			NodeValues: []interface{}{2, 3, 4},
			inData:     1,
		},
		"When list has no value": {
			NodeValues: []interface{}{},
			inData:     1,
		},
	}

	for k, v := range pattern {
		list := newLinkedList(v.NodeValues)

		t.Run(k, func(t *testing.T) {
			list.Remove(v.inData)

			actual := list.Head
			for actual != nil && actual.Next != nil {
				if actual.Data == v.inData {
					t.Errorf("Expected {%v} is deleted, but actual member: {%v}", v.inData, actual)
				}

				actual = actual.Next
			}
		})
	}
}

func TestReverse(t *testing.T) {
	pattern := map[string]struct {
		NodeValues  []interface{}
		expectOrder []interface{}
	}{
		"When numbers are ordered by desc": {
			NodeValues:  []interface{}{1, 2, 3, 4, 5},
			expectOrder: []interface{}{5, 4, 3, 2, 1},
		},
		"When chars are ordered by desc": {
			NodeValues:  []interface{}{"あ", "い", "う", "え", "お"},
			expectOrder: []interface{}{"お", "え", "う", "い", "あ"},
		},
	}

	for k, v := range pattern {
		list := newLinkedList(v.NodeValues)

		t.Run(k, func(t *testing.T) {
			list.Reverse()
			actual := list.Head

			for i := 0; i < len(v.expectOrder); i++ {
				if actual.Data != v.expectOrder[i] {
					t.Errorf("Expected value is %v, but actual is %v", v.expectOrder[i], actual)
				}

				actual = actual.Next
			}
		})
	}
}

func newLinkedList(nodeValues []interface{}) *LinkedList {
	list := new(LinkedList)

	for i := range nodeValues {
		newNode := &Node{Data: nodeValues[i]}

		if list.Head == nil {
			list.Head = newNode
			continue
		}

		lastNode := list.Head
		for lastNode.Next != nil {
			lastNode = lastNode.Next
		}
		lastNode.Next = newNode
	}

	return list
}
