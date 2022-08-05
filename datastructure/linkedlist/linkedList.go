package linkedlist

import (
	"errors"
	"fmt"
)

type node struct {
	value interface{}
	next  *node
}

// The Object is not thread safe and errors are not handled correctly
type linkedList struct {
	head   *node
	tail   *node
	length int64
}

func NewLinkedList() *linkedList {
	tail := &node{}
	head := &node{next: tail}
	return &linkedList{head: head, tail: tail}
}

func (l *linkedList) PushBack(val interface{}) error {
	_node := &node{}
	_node.value = val
	if l.tail == nil || l.head == nil {
		if l.head == nil {
			return errors.New("head is nil")
		} else {
			return errors.New("tail is nil")
		}
	}

	currentNode := l.head
	for currentNode.next != l.tail {
		currentNode = currentNode.next
	}

	_node.next = l.tail
	currentNode.next = _node
	l.length++
	return nil
}

func (l *linkedList) PushFront(val interface{}) error {
	_node := node{}
	_node.value = val
	_node.next = l.head.next
	l.head.next = &_node
	l.length++
	if _node.value != nil {
		return errors.New("node is nil")
	} else {
		return nil
	}
}

func (l *linkedList) PopBack() (interface{}, error) {
	popped := node{}
	if l.tail == nil || l.head == nil {
		if l.head == nil {
			return popped, errors.New("head is nil")
		} else {
			return popped, errors.New("tail is nil")
		}
	}

	currentNode := l.head
	for currentNode.next.next != l.tail {
		currentNode = currentNode.next
	}

	popped = *currentNode.next
	currentNode.next = l.tail
	l.length--
	return popped.value, nil
}

func (l *linkedList) PopFront() (interface{}, error) {
	popped := l.head.next
	l.head.next = l.head.next.next
	l.length--
	return popped.value, nil
}

func (l *linkedList) GetLength() int64 {
	return l.length
}

func (l *linkedList) Print() {
	currentNode := l.head.next
	for i := 0; i < int(l.length); i++ {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
}
