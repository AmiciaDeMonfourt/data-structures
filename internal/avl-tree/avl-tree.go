package avltree

type avltree struct {
	root *node
}

func New() *avltree {
	return &avltree{
		root: nil,
	}
}

func (t *avltree) Insert(key int, data any) {
	t.root = t.root.insert(key, data)
}
