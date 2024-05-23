package avltree

// AVL Tree leaf struct
type node struct {
	pair
	height int
	left   *node
	right  *node
}

// node information struct
type pair struct {
	data any
	key  int
}

func newNode(data any, key int) *node {
	return &node{
		pair: pair{
			data: data,
			key:  key,
		},
		height: 0,
		left:   nil,
		right:  nil,
	}
}

func (n *node) insert(key int, data any) *node {
	if n == nil {
		return newNode(data, key)
	}

	if key > n.key {
		n.right = n.right.insert(key, data)

	} else if key < n.key {
		n.left = n.left.insert(key, data)

	} else {
		n.data = data
	}

	return n.balanced()
}

func (n *node) balanced() *node {
	return n
}
