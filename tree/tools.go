package tree

import (
	"fmt"
)

// // PrettyPrint prints a binary tree in a nice, human-readable format.
// func PrettyPrint(root *Node) {
// 	if root == nil {
// 		fmt.Println("Empty Tree")
// 		return
// 	}
// 	prettyPrint(root, 0)
// }

// func prettyPrint(node *Node, level int) {
// 	if node == nil {
// 		return
// 	}

// 	prettyPrint(node.right, level+1)

// 	if level == 0 {
// 		fmt.Println(node.data)
// 	} else {
// 		for i := 1; i < level; i++ {
// 			fmt.Print("\t")
// 		}
// 		fmt.Printf("|---->%d\n", node.data)
// 	}

// 	prettyPrint(node.left, level+1)
// }

// PrettyPrint prints the binary tree in a clear, human-readable format.
func PrettyPrint(root *Node) {
	if root == nil {
		fmt.Println("The tree is empty.")
		return
	}
	printTree(root, "", true)
}

func printTree(node *Node, prefix string, isRight bool) {
	if node == nil {
		return
	}

	// Print the right subtree first so it appears on top
	if node.right != nil {
		printTree(node.right, prefix+"    ", true)
	}

	// Print current node
	connector := "└──" // For left child
	if isRight {
		connector = "┌──" // For right child
	}

	fmt.Printf("%s%s %d\n", prefix, connector, node.data)

	// Print the left subtree
	if node.left != nil {
		printTree(node.left, prefix+"    ", false)
	}
}
