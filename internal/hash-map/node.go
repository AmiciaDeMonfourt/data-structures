package hmap

type Node struct {
	key  string
	val  interface{}
	next *Node
}

func NewNode(key string, val interface{}) *Node {
	return &Node{
		key:  key,
		val:  val,
		next: nil,
	}
}
