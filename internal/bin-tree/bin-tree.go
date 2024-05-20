package bintree

import "fmt"

type binaryTree struct {
	root   *node
	height int
}

func New() *binaryTree {
	return &binaryTree{
		root: nil,
	}
}

func (t *binaryTree) Insert(data interface{}, key int) {
	newNode := newNode(data, key)
	cursor := t.root
	height := 0

	if cursor == nil {
		t.root = newNode
		return

	} else {

		for {
			height++
			if height > t.height {
				t.height = height
			}

			if key > cursor.Key {
				if cursor.Right == nil {
					cursor.Right = newNode
					return

				} else if key > cursor.Right.Key {
					newNode.Left = cursor.Right
					cursor.Right = newNode
					return
				}

				cursor = cursor.Right
				continue

			} else if key < cursor.Key {
				if cursor.Left == nil {
					cursor.Left = newNode
					return

				} else if key > cursor.Left.Key {
					newNode.Left = cursor.Left
					cursor.Left = newNode
					return
				}

				cursor = cursor.Left
				continue

			} else {
				cursor.Data = data
				return
			}
		}
	}
}

// Remove item from tree. Two cases:
//   if deleting item has no right child element
//		 replace the deleted item with the left child
//	 else
//	     find the minimum element in right subtree and replace it with the deleted element
func (t *binaryTree) Delete(key int) {
	if t.root == nil {
		panic("tree is empty, nothing to delete")
	}

}

func (t *binaryTree) Search(key int) interface{} {
	return t.findNodeByKey(key).Data
}

func (t *binaryTree) findParentByKey(key int) *node {
	child := t.root
	parent := t.root

	for child != nil {
		if key == child.Key {
			return parent

		} else if key < child.Key {
			parent = child
			child = child.Left
			continue

		} else if key > child.Key {
			parent = child
			child = child.Right
			continue
		}
	}

	panic(fmt.Sprintf("There is no record with this key: %d", key))
}

func (t *binaryTree) findNodeByKey(key int) *node {
	cursor := t.root

	for cursor != nil {
		if key == cursor.Key {
			return cursor

		} else if key < cursor.Key {
			cursor = cursor.Left
			continue

		} else if key > cursor.Key {
			cursor = cursor.Right
			continue
		}
	}

	panic(fmt.Sprintf("There is no record with this key: %d", key))
}

// pre order traverse, using recursive descent
func (t *binaryTree) TraverseTree() {
	var descent func(*node)

	descent = func(cursor *node) {
		if cursor == nil {
			return
		}
		fmt.Printf("%d ", cursor.Key)
		descent(cursor.Left)
		descent(cursor.Right)
	}

	descent(t.root)
}
