package tree

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func NewSimpleTree() *Node {
	var data int
	fmt.Printf("Enter Data for Root: ")
	fmt.Scanf("%d", &data)
	root := NewNode(data)
	buildSimpleTree(root)
	return root
}

func buildSimpleTree(root *Node) {
	var (
		yes  bool
		data int
	)

	yes = getChoice(fmt.Sprintf("Want to enter to the left of %d? (yes or no): ", root.data))
	if yes {
		fmt.Print("Enter Data: ")
		fmt.Scanf("%d", &data)
		node := NewNode(data)
		root.left = node
		buildSimpleTree(root.left)
	}

	yes = getChoice(fmt.Sprintf("Want to enter to the right of %d? (yes or no): ", root.data))
	if yes {
		fmt.Print("Enter Data: ")
		fmt.Scanf("%d", &data)
		node := NewNode(data)
		root.right = node
		buildSimpleTree(root.right)
	}
}

func getChoice(text string) bool {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(text)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	switch input {
	case "yes", "true", "1":
		return true
	case "no", "false", "0":
		return false
	default:
		panic(fmt.Sprintf("invalid input %v", input))
	}
}
