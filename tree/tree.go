package tree

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func NewSimpleTree() *Node {
	var data int
	fmt.Printf("Enter Data for Root: ")
	fmt.Scanf("%d", &data)
	root := NewNode(data)
	buildSimpleTree(root)
	root.height = int(math.Max(float64(root.left.Height()), float64(root.right.Height())) + 1)
	return root
}

func buildSimpleTree(node *Node) {
	var (
		yes  bool
		data int
	)

	yes = getChoice(fmt.Sprintf("Want to enter to the left of %d? (yes or no): ", node.data))
	if yes {
		fmt.Print("Enter Data: ")
		fmt.Scanf("%d", &data)
		node := NewNode(data)
		node.left = node
		buildSimpleTree(node.left)
		node.height = int(math.Max(float64(node.left.Height()), float64(node.right.Height())) + 1)
	}

	yes = getChoice(fmt.Sprintf("Want to enter to the right of %d? (yes or no): ", node.data))
	if yes {
		fmt.Print("Enter Data: ")
		fmt.Scanf("%d", &data)
		node := NewNode(data)
		node.right = node
		buildSimpleTree(node.right)
		node.height = int(math.Max(float64(node.left.Height()), float64(node.right.Height())) + 1)
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
