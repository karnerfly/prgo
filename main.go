package main

import (
	"main/tree"
)

func main() {
	// data := []int{1, 2, 3, 6, 7, 8, 11, 12, 23}
	// // data := []int{2, 4, 3, 5, 1}
	// i := search.BinarySearch(data, 11)
	// fmt.Println(i)

	root := tree.NewSimpleTree()
	tree.PrettyPrint(root)
}
