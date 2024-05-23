package bintree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree_Insert(t *testing.T) {
	tree := getTestBinaryTree()
	assert.Equal(
		t,
		[]int{6, 1, 10, 8, 7, 9, 11},
		tree.GetKeysInOrder(),
	)
}

func TestBinaryTree_GetNodesInOrder(t *testing.T) {
	tree := getTestBinaryTree()
	assert.Equal(
		t,
		[]pair{{6, "6"}, {1, "1"}, {10, "10"}, {8, "8"}, {7, "7"}, {9, "9"}, {11, "11"}},
		tree.GetNodesInOrder(),
	)

}

func TestBinaryTree_SearchByKey(t *testing.T) {
	tree := getTestBinaryTree()
	testCases := []struct {
		pair
		isValid bool
	}{
		{
			pair: pair{
				key:  0,
				data: "0",
			},
			isValid: false,
		},
		{
			pair: pair{
				key:  1,
				data: "1",
			},
			isValid: true,
		},
		{
			pair: pair{
				key:  6,
				data: "6",
			},
			isValid: true,
		},
	}

	for _, tc := range testCases {
		if tc.isValid {
			assert.Equal(t, tc.pair, tree.SearchByKey(tc.key))
		} else {
			assert.Equal(t, pair{}, tree.SearchByKey(tc.key))
		}
	}
}

func TestBinaryTree_SearchByKeys(t *testing.T) {
	tree := getTestBinaryTree()

	nodes := tree.SearchNodesByKeys([]int{0, 1, 9, -24})
	assert.Equal(t, []pair{{1, "1"}, {9, "9"}}, nodes)

	nodes = tree.SearchNodesByKeys([]int{-24})
	assert.Equal(t, []pair{}, nodes)
}

func TestBinaryTree_SearchNodesByData(t *testing.T) {
	tree := getTestBinaryTree()

	nodes := tree.SearchNodesByData([]any{"6", "11", "-412"})

	assert.Equal(t, []pair{{6, "6"}, {11, "11"}}, nodes)
}

func TestBinaryTree_SearchNodesByKeys(t *testing.T) {
	tree := getTestBinaryTree()

	nodes := tree.SearchNodesByKeys([]int{6, 9, -7777})

	assert.Equal(t, []pair{{6, "6"}, {9, "9"}}, nodes)
}

func TestBinaryTree_Delete(t *testing.T) {
	tree := getTestBinaryTree()

	// delete node without childs
	tree.Delete(7)
	assert.Equal(t, []int{6, 1, 10, 8, 9, 11}, tree.GetKeysInOrder())

	// deleete node with only right child
	tree.Delete(10)
	assert.Equal(t, []int{6, 1, 11, 8, 9}, tree.GetKeysInOrder())

	// delete node with 2 childs
	tree.Delete(8)
	assert.Equal(t, []int{6, 1, 11, 9}, tree.GetKeysInOrder())

	tree.Insert("t", 12)

	// delete node with only left child
	tree.Delete(11)
	assert.Equal(t, []int{6, 1, 12, 9}, tree.GetKeysInOrder())
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
