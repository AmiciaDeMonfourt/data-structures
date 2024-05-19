package list

type Node struct {
	Data interface{}
	Next *Node
	Prev *Node
}

func newNode(data interface{}) *Node {
	return &Node{
		Data: data,
		Next: nil,
		Prev: nil,
	}

}
