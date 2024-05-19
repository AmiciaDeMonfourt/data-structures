package list

import (
	"fmt"
	"log"
)

// doubly linked list
type List struct {
	head *Node
	tail *Node
	size uint
}

func NewList() *List {
	return &List{
		head: nil,
		tail: nil,
		size: 0,
	}
}

func (l *List) PushAfterNode(data interface{}, index uint) {
	node := newNode(data)

	if l.size == 0 {
		l.head = node
		l.tail = node

	} else {
		cursor := l.iterateAt(index)

		// if cursor is tail
		if cursor == l.tail {
			l.tail = node

		} else {
			// is there is addiion between elements
			cursor.Next.Prev = node
			node.Next = cursor.Next
		}

		// update pointers
		cursor.Next = node
		node.Prev = cursor
	}

	l.size++
}

func (l *List) PushBeforeNode(data interface{}, index uint) {
	node := newNode(data)

	// if list is empty - replace just replace new node with head
	if l.size == 0 {
		l.head = node
		l.tail = node

	} else {
		cursor := l.iterateAt(index)

		// if cursor is head
		if cursor == l.head {
			l.head = node

		} else {
			// is there is addiion between elements
			cursor.Prev.Next = node
			node.Prev = cursor.Prev
		}

		// update pointers
		cursor.Prev = node
		node.Next = cursor
	}

	l.size++
}

func (l *List) PushFront(data interface{}) {
	node := newNode(data)

	// if list is empty - replace just replace new node with head
	if l.size == 0 {
		l.head = node
		l.tail = node

	} else {
		// if list isn't empty, update pointers
		l.head.Prev = node
		node.Next = l.head
		l.head = node
	}

	l.size++
}

func (l *List) PushBack(data interface{}) {
	node := newNode(data)

	// if list is empty - replace just replace new node with head
	if l.size == 0 {
		l.head = node
		l.tail = node

	} else {
		// if list isn't empty, update pointers
		l.tail.Next = node
		node.Prev = l.tail
		l.tail = node
	}

	l.size++
}

func (l *List) GetElem(index uint) *Node {
	return l.iterateAt(index)
}

func (l *List) iterateAt(index uint) *Node {
	if index > l.size {
		log.Fatal("List.iterateAt: index > list size")
	}

	if index == 0 {
		return l.head
	}

	node := l.head
	if index < l.size/2 {
		for i := 0; i < int(index); i++ {
			node = node.Next
		}

	} else {
		node = l.tail
		for i := 0; i < int(l.size-index-1); i++ {
			node = node.Prev
		}
	}

	return node
}

func (l *List) Print() {
	it := l.head

	for i := range l.size {
		fmt.Printf("%d: %v\n", i, it.Data)
		it = it.Next
	}
}

func (l *List) GetSize() uint {
	return l.size
}
