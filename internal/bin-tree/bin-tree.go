package bintree

type BinaryTree struct {
	root *node
}

func New() *BinaryTree {
	return &BinaryTree{
		root: nil,
	}
}

func (t *BinaryTree) Insert(data any, key int) {
	t.root = t.root.insert(data, key)
}

func (t *BinaryTree) Delete(key int) {
	t.root = t.root.remove(key)
}

func (t *BinaryTree) SearchByKey(key int) *node {
	return t.root.searchByKey(key)
}

func (t *BinaryTree) SearchNodesByData(data []any) []pair {
	return t.root.searchNodesByData(data)
}

func (t *BinaryTree) SearchNodesByKeys(keys []int) []pair {
	return t.root.searchNodesByKeys(keys)
}

func (t *BinaryTree) GetNodesInOrder() []node {
	return t.root.nodeTraversal()
}

func (t *BinaryTree) GetKeysInOrder() []int {
	return t.root.keyTraversal()
}
