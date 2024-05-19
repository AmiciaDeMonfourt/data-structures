package hmap

import "fmt"

// singly linked list
type List struct {
	head *Node
}

func (l *List) Push(key string, val interface{}) {
	if key == "" {
		panic("hmap.insert(): key isn't can be empty!")
	}

	if l.head == nil {
		l.head = NewNode(key, val)

	} else {
		cursor := l.head

		// going through every element until the last
		for cursor.next != nil {
			// if map has node with this key
			if cursor.key == key {
				cursor.val = val
				return
			}

			cursor = cursor.next
		}

		// checking last element
		if cursor.key == key {
			cursor.val = val
			return
		}

		cursor.next = NewNode(key, val)
	}
}

// remove an element from the list
func (l *List) Pull(key string) {
	cursor := l.head
	if cursor.key == key {
		l.head = l.head.next
		return
	}

	for cursor.next != nil {

		if cursor.next.key == key {
			cursor.next = cursor.next.next
			return
		}

		cursor = cursor.next
	}

	panic("hmap.delete(): key is not found")
}

func (l *List) Search(key string) interface{} {
	cursor := l.head

	for cursor != nil {
		if cursor.key == key {
			// fmt.Println(cursor.val)
			return cursor.val
		}
		cursor = cursor.next
	}

	panic(fmt.Sprintf("unknown record, key: %s", key))
}

func (l *List) Print() (items []string) {
	cursor := l.head
	var str string

	for cursor != nil {
		str = fmt.Sprintf("{%s : %v}", cursor.key, cursor.val)

		if cursor = cursor.next; cursor != nil {
			str += ", "
		}

		items = append(items, str)
	}

	return items
}
