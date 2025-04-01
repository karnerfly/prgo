package main

import (
	"fmt"
	"main/conv"
)

func main() {
	// str := "Rony Chutiya, bokachoda!!!!"
	// h := huffman.New(str)
	// n := h.Encode()

	// fmt.Printf("Original size %d, Compressed size %d\n", len(str)*8, len(n))

	fmt.Println(conv.CheckBalancedExpr("[]()"))
}
