package tree

import (
	"fmt"
)

// PrettyPrint prints a binary tree in a nice, human-readable format.
func PrettyPrint(root *Node) {
	if root == nil {
		fmt.Println("Empty Tree")
		return
	}
	prettyPrint(root, 0)
}

func prettyPrint(node *Node, level int) {
	if node == nil {
		return
	}

	prettyPrint(node.right, level+1)

	if level == 0 {
		fmt.Println(node.data)
	} else {
		for i := 0; i < level; i++ {
			fmt.Print("\t")
		}
		fmt.Println("----->", node.data)
	}

	prettyPrint(node.left, level+1)
}
