package tree

type Node struct {
	data   int
	height int
	left   *Node
	right  *Node
}

func NewNode(data int) *Node {
	return &Node{data: data}
}

func (n *Node) Height() int {
	if n == nil {
		return -1
	}
	return n.height
}

func (n *Node) Data() int {
	return n.data
}
