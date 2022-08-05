package queue

import "test3/dataStructure/linkedlist"

type LinkedListQueue interface {
	PushFront(val interface{}) error
	PopBack() (interface{}, error)
	Print()
	GetLength() int64
}

type Queue struct {
	LinkedListQueue
}

func NewStack() LinkedListQueue {
	return &Queue{linkedlist.NewLinkedList()}
}
