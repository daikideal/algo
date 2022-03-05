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
