package stack

import "test3/dataStructure/linkedlist"

type LinkedListStack interface {
	PushFront(val interface{}) error
	PopFront() (interface{}, error)
	Print()
	GetLength() int64
}

func NewStack() LinkedListStack {
	return linkedlist.NewLinkedList()
}
