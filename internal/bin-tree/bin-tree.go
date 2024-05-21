package bintree

import "fmt"

type BinaryTree struct {
	root   *node
	height int
}

func New() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

// Insert new item in binary tree
// Insert into an empty place, that finded by rule:
// if new node key > current subtree root key -> descent to the right. Else descent to the left
func (t *BinaryTree) Insert(data interface{}, key int) {
	newNode := newNode(data, key)
	cursor := t.root
	height := 0

	if cursor == nil {
		t.root = newNode
		return

	} else {
		for {
			height++
			// calculate binary tree height
			if height > t.height {
				t.height = height
			}

			if key > cursor.Key {
				// if right element is free - add new element
				if cursor.Right == nil {
					cursor.Right = newNode
					return
				}
				// else descent down
				cursor = cursor.Right
				continue

			} else if key < cursor.Key {
				// is there left element is free - add new element on this place
				if cursor.Left == nil {
					cursor.Left = newNode
					return
				}
				// else descent down
				cursor = cursor.Left
				continue

				// if tree has a node with equal key
			} else {
				cursor.Data = data
				return
			}
		}
	}
}

// Remove item from tree. Four cases by element to be deleted childs:
// 1)left and right nil
// 2)left nil
// 3)right nil
// 4)left and right child not nil
func (t *BinaryTree) Delete(key int) {
	if t.root == nil {
		panic("tree is empty, nothing to delete")
	}

	rmNode := t.findNodeByKey(key)
	rmNodeParent := t.findParentByKey(key)

	// Case where record to be deleted has only the left child.
	// Just replace the right node and the node to be deleted
	if rmNode.Left == nil && rmNode.Right == nil {
		if rmNodeParent.Left == rmNode {
			rmNodeParent.Left = nil
		} else {
			rmNodeParent.Right = nil
		}

		// Case where record to be deleted has only the left child.
		// Just replace the right node and the node to be deleted
	} else if rmNode.Right == nil {
		if rmNodeParent.Left == rmNode {
			rmNodeParent.Left = rmNode.Left
		} else {
			rmNodeParent.Right = rmNode.Left
		}
		return

		// Case where record to be deleted has only the right child
		// 1)find minimal element and his parent in the right subtree,
		// 2)replace min and the record to be deleted
		// 3)update the minimal element and the record to be deleted parent's pointers
	} else if rmNode.Left == nil {
		min, minParent := t.findMinumum(rmNode.Right)
		minParent.Left = nil

		if rmNodeParent.Left == rmNode {
			rmNodeParent.Left = min
		} else {
			rmNodeParent.Right = min
		}
		return

		// Case where: record to be deleted has only the right child
		// 1)find minimal element and his parent in the right subtree
		// 2)replace min and the record to be deleted
		// 3)update the minimal element and the record to be deleted parent's pointers
	} else {
		min, minParent := t.findMinumum(rmNode.Right)
		minParent.Left = nil

		if rmNodeParent.Left == rmNode {
			rmNodeParent.Left = min
		} else {
			rmNodeParent.Right = min
		}

		// update pointers at the new subtree root
		min.Left = rmNode.Left
		if rmNode.Right != min {
			min.Right = rmNode.Right
		}

		return
	}

}

func (t *BinaryTree) findMinumum(subtreeRoot *node) (min *node, hisParent *node) {
	if subtreeRoot == nil {
		return nil, nil
	}

	parent := subtreeRoot
	for subtreeRoot.Left != nil {
		parent = subtreeRoot
		subtreeRoot = subtreeRoot.Left
	}

	return subtreeRoot, parent
}

func (t *BinaryTree) Search(key int) interface{} {
	return t.findNodeByKey(key).Data
}

func (t *BinaryTree) findParentByKey(key int) *node {
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

func (t *BinaryTree) findNodeByKey(key int) *node {
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
func (t *BinaryTree) TraverseTree() []interface{} {
	var descent func(*node)
	keys := make([]interface{}, 0)

	descent = func(cursor *node) {
		if cursor == nil {
			return
		}
		keys = append(keys, cursor.Key)

		descent(cursor.Left)
		descent(cursor.Right)
	}

	descent(t.root)
	return keys
}

func (t *BinaryTree) GetHeight() int {
	return t.height
}
