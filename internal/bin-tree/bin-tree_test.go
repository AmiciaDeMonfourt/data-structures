package bintree

import (
	"fmt"
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

	fmt.Println(tree.GetKeysInOrder())

	assert.Equal(
		t,
		[]interface{}{6, 1, 0, -1, 2, 8, 9, 18},
		tree.GetKeysInOrder(),
	)
}

func TestBinaryTree_SearchByKey(t *testing.T) {
	tree := New()

	assert.Panics(t, func() {
		tree.SearchByKey(0)
	})
	tree.Insert("6", 6)
	tree.Insert("8", 8)
	tree.Insert("1", 1)
	tree.Insert("9", 9)

	assert.Equal(t, "1", tree.SearchByKey(1))

	assert.Equal(t, "9", tree.SearchByKey(9))

	assert.NotEqual(t, "2", 6)

	assert.Panics(t, func() {
		tree.SearchByKey(18)
	})
}

func TestBinaryTree_SearchByKeys(t *testing.T) {
	tree := getTestBinaryTree()

	nodes := tree.SearchNodesByKeys([]int{1, 7, 11})

	fmt.Println(nodes)
}

func TestBinaryTree_Delete(t *testing.T) {
	tree := getTestBinaryTree()

	// delete node without childs
	tree.Delete(7)
	assert.Equal(t, []any{6, 1, 10, 8, 9, 11}, tree.GetKeysInOrder())

	// deleete node with only right child
	tree.Delete(10)
	assert.Equal(t, []any{6, 1, 11, 8, 9}, tree.GetKeysInOrder())

	// delete node with 2 childs
	tree.Delete(8)
	assert.Equal(t, []any{6, 1, 11, 9}, tree.GetKeysInOrder())

	tree.Insert("t", 12)

	// delete node with only left child
	tree.Delete(11)
	assert.Equal(t, []any{6, 1, 12, 9}, tree.GetKeysInOrder())
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
