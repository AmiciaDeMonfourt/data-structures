package bintree

type node struct {
	Key   int
	Data  interface{}
	Left  *node
	Right *node

	Level int // delete
}

func newNode(data interface{}, key int) *node {
	return &node{
		Key:   key,
		Data:  data,
		Left:  nil,
		Right: nil,
		Level: 0,
	}
}
