package huffman

import (
	"main/heap"
)

type node struct {
	name  rune
	cost  int
	left  *node
	right *node
}

func newNode(name rune, cost int) *node {
	return &node{
		name: name,
		cost: cost,
	}
}

type Codec struct {
	data         string
	frequencyMap map[rune]int
	encoderMap   map[rune][]byte
	heap         *heap.Heap[[]*node, *node]
}

func New(data string) *Codec {
	return &Codec{
		data:         data,
		frequencyMap: make(map[rune]int),
		encoderMap:   make(map[rune][]byte),
		heap:         heap.New(make([]*node, 0), func(a, b *node) bool { return a.cost < b.cost }),
	}
}

func (c *Codec) Encode() []byte {
	if c.data == "" {
		return nil
	}

	// create frequency map of given text
	for _, r := range c.data {
		v, ok := c.frequencyMap[r]
		if !ok {
			c.frequencyMap[r] = 1
		} else {
			c.frequencyMap[r] = v + 1
		}
	}

	// for every entry in frequency map create a node with the character name and frequency as cost
	// insert into min heap
	for ch := range c.frequencyMap {
		node := newNode(ch, c.frequencyMap[ch])
		c.heap.Insert(node)
	}

	// until one element left in heap pop two elements merge them create new node with left as first and right as second and put it to the heap again
	for c.heap.Size() > 1 {
		first, _ := c.heap.Pop()
		second, _ := c.heap.Pop()
		n := newNode(0, first.cost+second.cost)
		n.left = first
		n.right = second
		c.heap.Insert(n)
	}

	// from the root traverse the tree
	root, _ := c.heap.Pop()
	c.traverse(root, make([]byte, 0))

	series := make([]byte, 0)
	for _, r := range c.data {
		series = append(series, c.encoderMap[r]...)
	}
	return series
}

// traverse, when goto left path add 0 to the edges and when goto right add 1 to the edges
func (c *Codec) traverse(n *node, edges []byte) {
	// if the node is a leaf node then add the character to the encoder map with its edge values
	if n.left == nil && n.right == nil {
		c.encoderMap[n.name] = append(c.encoderMap[n.name], edges...)
		return
	}

	c.traverse(n.left, append(edges, 0))
	c.traverse(n.right, append(edges, 1))
}

// func prettyPrint(n *node, level int) {
// 	if n == nil {
// 		return
// 	}

// 	prettyPrint(n.right, level+1)

// 	if level == 0 {
// 		fmt.Printf("(%q,%v)\n", n.name, n.cost)
// 	} else {
// 		for i := 0; i < level; i++ {
// 			fmt.Print("\t")
// 		}
// 		fmt.Printf("---->(%q,%v)\n", n.name, n.cost)
// 	}

// 	prettyPrint(n.left, level+1)
// }
