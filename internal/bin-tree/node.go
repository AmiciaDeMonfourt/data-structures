package bintree

// Binary tree leaf struct
type node struct {
	pair
	left  *node
	right *node
}

// node information structure
type pair struct {
	key  int
	data interface{}
}

// Node constructor
func newNode(data interface{}, key int) *node {
	return &node{
		pair: pair{
			key:  key,
			data: data,
		},
		left:  nil,
		right: nil,
	}
}

// Add a new node
func (n *node) insert(data interface{}, key int) *node {
	if n == nil {
		return newNode(data, key)
	}

	// binary searching
	if key > n.key {
		n.right = n.right.insert(data, key)

	} else if key < n.key {
		n.left = n.left.insert(data, key)

	} else {
		n.data = data
	}

	return n
}

// Remove a node by key
func (n *node) remove(key int) *node {
	if n == nil {
		return nil
	}

	// look for the desired node
	if key > n.key {
		n.right = n.right.remove(key)

	} else if key < n.key {
		n.left = n.left.remove(key)

	} else {
		// when node is found, we have four cases:
		// node has no childrens - just delete it
		if n.left == nil && n.right == nil {
			n = nil
		} else if n.right == nil {
			// node has only left children
			// replace node to delete with the left children
			n = n.left
		} else if n.left == nil {
			// node has only right children
			// replace node to delete with the right children
			n = n.right
		} else {
			//find minimal element in the right subtree
			min := n.right.findMin()
			// replace min and the pair to be deleted values
			n.key = min.key
			n.data = min.data
			// remove min element
			n.right = n.right.remove(min.key)
		}
	}
	return n
}

// Find minimum relative to current root
func (n *node) findMin() *node {
	if n.left != nil {
		return n.left.findMin()
	}
	return n
}

// Binary searching for a node by key
func (n *node) searchByKey(key int) *node {
	if n == nil {
		return nil
	}

	if key > n.key {
		return n.right.searchByKey(key)
	} else if key < n.key {
		return n.left.searchByKey(key)
	}

	return n
}

// Binary searching for group nodes by their keys
func (n *node) searchNodesByKeys(keys []int) []pair {
	if n == nil || keys == nil || len(keys) == 0 {
		return nil
	}
	// make a key set from the entry slice
	keysMap := make(map[int]struct{}, len(keys))
	for idx := range keys {
		keysMap[keys[idx]] = struct{}{}
	}

	pairs := make([]pair, 0)
	// adds a copy of the node pair, if the keys match
	callback := func(node *node) {
		if _, ok := keysMap[node.key]; ok {
			pairs = append(pairs, node.pair)
		}
	}
	n.traversalWithAction(callback)
	return pairs
}

func (n *node) searchNodesByData(data []interface{}) []pair {
	if n == nil || data == nil {
		return nil
	}
	// makes a data set from the entry slice
	dataMap := make(map[any]struct{}, 0)
	for idx := range data {
		dataMap[data[idx]] = struct{}{}
	}

	pairs := make([]pair, 0)
	// adds a copy ofs the node pair, if the data match
	callback := func(node *node) {
		if _, ok := dataMap[node.data]; ok {
			pairs = append(pairs, node.pair)
		}
	}
	n.traversalWithAction(callback)
	return pairs
}

// Return all the nodes relatives to the current root
func (n *node) nodeTraversal() []pair {
	nodes := make([]pair, 0)
	// adds copy of the node
	callback := func(node *node) {
		nodes = append(nodes, node.pair)
	}
	n.traversalWithAction(callback)
	return nodes
}

// Return keys of all elements relatieves to the current root
func (n *node) keyTraversal() []int {
	nodes := make([]int, 0)
	// adds key of the node to slice
	callback := func(node *node) {
		nodes = append(nodes, node.key)
	}
	n.traversalWithAction(callback)
	return nodes
}

// Tree descent with a function callback for each node
func (n *node) traversalWithAction(callback func(*node)) {
	if n == nil {
		return
	}
	// some action
	callback(n)
	// descent to down
	n.left.traversalWithAction(callback)
	n.right.traversalWithAction(callback)
}
