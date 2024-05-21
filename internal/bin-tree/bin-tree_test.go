package bintree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_Insert(t *testing.T) {
	tree := New()

	tree.Insert("data1", 6)
	tree.Insert("data2", 8)
	tree.Insert("data3", 1)
	tree.Insert("data4", 2)
	tree.Insert("data5", 9)
	tree.Insert("data5", 0)
	tree.Insert("data4", -1)
	tree.Insert("data4", 18)

	assert.Equal(
		t,
		[]interface{}{6, 1, 0, -1, 2, 8, 9, 18},
		tree.TraverseTree(),
	)
}

func TestBinaryTree_Search(t *testing.T) {
	tree := New()

	assert.Panics(t, func() {
		tree.Search(0)
	})
	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("1", 1)
	tree.Insert("9", 9)

	assert.Equal(t, "1", tree.Search(1))

	assert.Equal(t, "9", tree.Search(9))

	assert.NotEqual(t, "2", 6)

	assert.Panics(t, func() {
		tree.Search(18)
	})
}

func TestBinaryTree_Delete(t *testing.T) {
	tree := getTestBinaryTree()

	// delete node without childs
	tree.Delete(7)
	assert.Equal(t, []any{6, 1, 10, 8, 9, 11}, tree.TraverseTree())

	// deleete node with only right child
	tree.Delete(10)
	assert.Equal(t, []any{6, 1, 11, 8, 9}, tree.TraverseTree())

	// delete node with 2 childs
	tree.Delete(8)
	assert.Equal(t, []any{6, 1, 11, 9}, tree.TraverseTree())

	// delete node with only left child
	tree.Delete(11)
	assert.Equal(t, []any{6, 1, 9}, tree.TraverseTree())
}

func TestBinaryTree_FindMinimum(t *testing.T) {
	tree := getTestBinaryTree()

	// minimal in tree
	min, minParent := tree.findMinumum(tree.findNodeByKey(6))
	assert.Equal(t, 1, min.Key)
	assert.Equal(t, 6, minParent.Key)

	// minimal in right root subtree
	min, minParent = tree.findMinumum(tree.findNodeByKey(10))
	assert.Equal(t, 7, min.Key)
	assert.Equal(t, 8, minParent.Key)

	// minimal in subtree, but substree has one node
	min, minParent = tree.findMinumum(tree.findNodeByKey(11))
	assert.Equal(t, 11, min.Key)
	assert.Equal(t, 11, minParent.Key)
}

func getTestBinaryTree() *BinaryTree {
	tree := New()

	tree.Insert("6", 6)
	tree.Insert("1", 1)
	tree.Insert("10", 10)
	tree.Insert("8", 8)
	tree.Insert("7", 7)
	tree.Insert("9", 9)
	tree.Insert("11", 11)

	return tree
}
