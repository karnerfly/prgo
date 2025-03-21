package tree

import (
	"fmt"
	"strconv"
)

// PrettyPrint prints a binary tree in a nice, human-readable format.
func PrettyPrint(root *Node) {
	if root == nil {
		return
	}
	maxDepth := maxDepth(root)
	levels := make([][]string, maxDepth+1)
	for i := range levels {
		levels[i] = make([]string, 1<<(i+1)-1)
	}
	fillLevels(root, levels, 0, 0)
	for _, level := range levels {
		for _, value := range level {
			fmt.Printf("%-5s", value)
		}
		fmt.Println()
	}
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.left)
	rightDepth := maxDepth(root.right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func fillLevels(node *Node, levels [][]string, level int, index int) {
	if node == nil {
		return
	}
	levels[level][index] = strconv.Itoa(node.data)
	fillLevels(node.left, levels, level+1, 2*index+1)
	fillLevels(node.right, levels, level+1, 2*index+2)
}
