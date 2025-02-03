package linkedlist

import "errors"

type Node struct {
	value int
	next  *Node
}

type List struct {
	head *Node
	size int
}

func New(a []int) *List {
	list := &List{}
	for _, i := range a {
		list.Push(i)
	}

	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(element int) {
	l.head = &Node{
		value: element,
		next:  l.head,
	}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size < 1 {
		return 0, errors.New("List is empty")
	}

	value := l.head.value
	l.head = l.head.next
	l.size--
	return value, nil
}

func (l *List) Array() []int {
	currentNode := l.head
	result := make([]int, l.size)
	for i := 0; i < l.size; i++ {
		result[l.size-i-1] = currentNode.value
		currentNode = currentNode.next
	}

	return result
}

func (l *List) Reverse() *List {
	currentNode := l.head
	list := &List{}
	for currentNode != nil {
		list.Push(currentNode.value)
		currentNode = currentNode.next
	}

	return list
}