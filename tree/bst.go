package tree

import (
	"math"
)

func NewBST(data []int) *Node {
	if data == nil {
		return nil
	}

	var root *Node
	for i := 0; i < len(data); i++ {
		root = buildBST(root, data[i])
	}
	root.height = int(math.Max(float64(root.left.Height()), float64(root.right.Height())) + 1)
	return root
}

func buildBST(node *Node, data int) *Node {
	if node == nil {
		return NewNode(data)
	}

	if data <= node.data {
		node.left = buildBST(node.left, data)
		node.height = int(math.Max(float64(node.left.Height()), float64(node.right.Height())) + 1)
	} else {
		node.right = buildBST(node.right, data)
		node.height = int(math.Max(float64(node.left.Height()), float64(node.right.Height())) + 1)
	}

	return node
}
