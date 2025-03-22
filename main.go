package main

import (
	"fmt"
	"main/huffman"
)

func main() {
	text := "Hello How are you? Hope so well doing good"
	h := huffman.New(text)
	series := h.Encode()
	fmt.Println("Original:", len(text)*8, "bytes")
	fmt.Println("Compressed:", len(series), "bytes")
}
