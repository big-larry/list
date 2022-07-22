package suckutils

type LinkedListNode[T any] struct {
	next  *LinkedListNode[T]
	Value T
}

type LinkedList[T any] struct {
	root   *LinkedListNode[T]
	last   *LinkedListNode[T]
	length uint32
}

func NewLinkedList[T any]() *LinkedList[T] {
	return &LinkedList[T]{}
}

func (l *LinkedList[T]) GetLength() uint32 {
	return l.length
}

func (l *LinkedList[T]) Add(item T) *LinkedListNode[T] {
	l.length++
	if l.root == nil {
		l.root = &LinkedListNode[T]{Value: item}
		l.last = l.root
		return l.root
	}
	l.last.next = &LinkedListNode[T]{Value: item}
	l.last = l.last.next
	return l.last
}

func (l *LinkedList[T]) GetFirst() *LinkedListNode[T] {
	return l.root
}

func (l *LinkedList[T]) GetLast() *LinkedListNode[T] {
	return l.last
}

func (l *LinkedListNode[T]) GetNext() *LinkedListNode[T] {
	return l.next
}

func (l *LinkedList[T]) Get(index uint32) *LinkedListNode[T] {
	if index >= l.length {
		return nil
	}
	node := l.root
	var i uint32
	for i = 1; i < index; i++ {
		node = node.next
	}
	return node
}
