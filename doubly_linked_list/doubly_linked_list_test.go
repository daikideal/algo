package doubly_linked_list

import (
	"testing"
)

func TestAppend(t *testing.T) {
	pattern := map[string]struct {
		setListValues []interface{}
		inData        interface{}
	}{
		"When list has no value, should set head of list": {
			setListValues: nil,
			inData:        1,
		},
		"When list has values, should set last": {
			setListValues: []interface{}{1, 2, 3, 4},
			inData:        5,
		},
	}

	for k, v := range pattern {
		list := newDoublyLinkedList(v.setListValues)

		t.Run(k, func(t *testing.T) {
			list.Append(v.inData)

			actual := list.Head
			for actual.Next != nil {
				actual = actual.Next
			}

			if actual.Data != v.inData {
				t.Errorf("Expected %v, but actual %v", v.inData, actual.Data)
			}
		})
	}
}

func TestInsert(t *testing.T) {
	pattern := map[string]struct {
		setListValues []interface{}
		inData        interface{}
	}{
		"When list has no value, should set head of list": {
			setListValues: nil,
			inData:        1,
		},
		"When list has values, should set head of list": {
			setListValues: []interface{}{1, 2, 3, 4},
			inData:        5,
		},
	}

	for k, v := range pattern {
		list := newDoublyLinkedList(v.setListValues)

		t.Run(k, func(t *testing.T) {
			list.Insert(v.inData)

			actual := list.Head
			if actual.Data != v.inData {
				t.Errorf("Expected %v, but actual %v", v.inData, actual.Data)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	pattern := map[string]struct {
		setListValue []interface{}
		inData       interface{}
	}{
		"When list head equals input value, should remove it": {
			setListValue: []interface{}{1, 2, 3, 4, 5},
			inData:       1,
		},
		"When list has input value, should remove it": {
			setListValue: []interface{}{1, 2, 3, 4, 5},
			inData:       3,
		},
		"When list does not have input value, should do nothing": {
			setListValue: []interface{}{1, 2, 3, 4, 5},
			inData:       6,
		},
	}

	for k, v := range pattern {
		list := newDoublyLinkedList(v.setListValue)

		t.Run(k, func(t *testing.T) {
			list.Remove(v.inData)

			actual := list.Head
			for actual.Next != nil {
				if actual.Data == v.inData {
					t.Errorf("%v is should not found", v.inData)
				}

				actual = actual.Next
			}
		})
	}
}

func TestReverse(t *testing.T) {
	pattern := map[string]struct {
		setListValue []interface{}
		expOrder     []interface{}
	}{
		"When numbers are ordered by desc": {
			setListValue: []interface{}{1, 2, 3, 4, 5},
			expOrder:     []interface{}{5, 4, 3, 2, 1},
		},
		"When chars are ordered by desc": {
			setListValue: []interface{}{"あ", "い", "う", "え", "お"},
			expOrder:     []interface{}{"お", "え", "う", "い", "あ"},
		},
	}

	for k, v := range pattern {
		list := newDoublyLinkedList(v.setListValue)

		t.Run(k, func(t *testing.T) {
			list.Reverse()
			actual := list.Head

			for i := 0; i < len(v.expOrder); i++ {
				if actual.Data != v.expOrder[i] {
					t.Errorf("Expected value is %v, but actual is %v", v.expOrder[i], actual)
				}

				actual = actual.Next
			}
		})
	}
}

func newDoublyLinkedList(values []interface{}) *DoublyLinkedList {
	list := new(DoublyLinkedList)

	for i := range values {
		newNode := &Node{Data: values[i]}

		if list.Head == nil {
			list.Head = newNode
			continue
		}

		currentNode := list.Head
		for currentNode.Next != nil {
			currentNode = currentNode.Next
		}
		currentNode.Next = newNode
		newNode.Prev = currentNode
	}

	return list
}
