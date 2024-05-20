package bintree_test

import (
	bintree "data-structures/internal/bin-tree"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_Insert(t *testing.T) {
	tree := bintree.New()

	tree.Insert("data1", 6)
	tree.Insert("data2", 8)
	tree.Insert("data3", 1)
	tree.Insert("data4", 2)
	tree.Insert("data5", 9)
	tree.Insert("data5", 0)
	tree.Insert("data4", -1)
	tree.Insert("data4", 18)

	tree.TraverseTree()
}

func TestBinaryTree_Search(t *testing.T) {
	tree := bintree.New()

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
	tree := bintree.New()

	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("1", 1)
	tree.Insert("9", 9)

	tree.Delete(6)
	tree.TraverseTree()
}
