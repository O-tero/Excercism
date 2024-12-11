package linkedlist

import (
	"errors"
)

// Node represents a node in the doubly linked list.
type Node struct {
	Value interface{}
	prev  *Node
	next  *Node
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}

type List struct {
	head *Node
	tail *Node
}

func NewList(elements ...interface{}) *List {
	list := &List{}
	for _, element := range elements {
		list.Push(element)
	}
	return list
}

func (l *List) First() *Node {
	return l.head
}

func (l *List) Last() *Node {
	return l.tail
}

func (l *List) Push(v interface{}) {
	newNode := &Node{Value: v}
	if l.tail == nil { 
		l.head = newNode
		l.tail = newNode
	} else {
		l.tail.next = newNode
		newNode.prev = l.tail
		l.tail = newNode
	}
}

func (l *List) Pop() (interface{}, error) {
	if l.tail == nil { 
		return nil, errors.New("cannot pop from an empty list")
	}
	value := l.tail.Value
	if l.head == l.tail { 
		l.head = nil
		l.tail = nil
	} else {
		l.tail = l.tail.prev
		l.tail.next = nil
	}
	return value, nil
}

func (l *List) Unshift(v interface{}) {
	newNode := &Node{Value: v}
	if l.head == nil { 
		l.head = newNode
		l.tail = newNode
	} else {
		l.head.prev = newNode
		newNode.next = l.head
		l.head = newNode
	}
}

func (l *List) Shift() (interface{}, error) {
	if l.head == nil {
		return nil, errors.New("cannot shift from an empty list")
	}
	value := l.head.Value
	if l.head == l.tail { 
		l.head = nil
		l.tail = nil
	} else {
		l.head = l.head.next
		l.head.prev = nil
	}
	return value, nil
}

func (l *List) Reverse() {
	current := l.head
	var prev *Node
	l.head, l.tail = l.tail, l.head
	for current != nil {
		next := current.next
		current.next = prev
		current.prev = next
		prev = current
		current = next
	}
}
