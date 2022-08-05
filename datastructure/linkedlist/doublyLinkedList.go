package linkedlist

import (
	"errors"
	"fmt"
)

type DoublyNode struct {
	value    interface{}
	next     *DoublyNode
	previous *DoublyNode
}

// The Object is not thread safe and errors are not handled correctly
type doublyLinkedList struct {
	head   *DoublyNode
	tail   *DoublyNode
	length int64
}

func NewDoublyLinkedList() doublyLinkedList {
	tail := &DoublyNode{}
	head := &DoublyNode{next: tail}
	tail.previous = head
	return doublyLinkedList{head: head, tail: tail}
}

func (dl *doublyLinkedList) PushBack(_node *DoublyNode) error {
	if dl.tail == nil || dl.head == nil {
		if dl.head == nil {
			return errors.New("head is nil")
		} else {
			return errors.New("tail is nil")
		}
	}
	_node.next = dl.tail
	_node.previous = dl.tail.previous
	dl.tail.previous.next = _node
	dl.tail.previous = _node

	dl.length++
	return nil
}

func (dl *doublyLinkedList) PushFront(_node DoublyNode) error {
	dl.head.next = &_node
	_node.next = dl.tail
	dl.length++
	if _node.value != nil {
		return errors.New("node is nil")
	} else {
		return nil
	}
}

func (dl *doublyLinkedList) PopBack() (DoublyNode, error) {
	popped := DoublyNode{}
	if dl.tail == nil || dl.head == nil {
		if dl.head == nil {
			return popped, errors.New("head is nil")
		} else {
			return popped, errors.New("tail is nil")
		}
	}
	popped = *dl.tail.previous
	dl.tail.previous.previous.next = dl.tail
	dl.tail.previous = dl.tail.previous.previous

	dl.length--
	return popped, nil
}

func (l *doublyLinkedList) PopFront() (DoublyNode, error) {
	popped := l.head.next
	l.head.next = l.head.next.next
	l.length--
	return *popped, nil
}

func (l *doublyLinkedList) GetLength() int64 {
	return l.length
}

func (l *doublyLinkedList) Print() {
	currentNode := l.head.next
	for i := 0; i < int(l.length); i++ {
		fmt.Println(currentNode.value)
		currentNode = currentNode.next
	}
}
