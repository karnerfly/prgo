package main

import (
	"fmt"
	"main/sort"
)

func main() {
	data := []int{19, 29, 32, 6, 75, 31, 17}
	sort.MergeSort(data)
	fmt.Println(data)
}
