package sort

// Classic Heap Sort Function. Time Complexity - O(nlog(n)), Space Complexity - O(1)
func HeapSort(data []int) {
	size := len(data)
	for i := size/2 - 1; i >= 0; i-- {
		heapify(data, size-1, i)
	}

	for i := size - 1; i >= 0; i-- {
		data[i] = removeFromHeap(data, i)
	}
}

func heapify(arr []int, n, i int) {
	largest := i
	left := i*2 + 1
	right := i*2 + 2

	if left <= n && arr[left] > arr[largest] {
		largest = left
	}

	if right <= n && arr[right] > arr[largest] {
		largest = right
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, n, largest)
	}
}

func removeFromHeap(arr []int, n int) int {
	elem := arr[0]
	arr[0] = arr[n]
	heapify(arr, n, 0)
	return elem
}

/*  How Implemented?

1. build the heap from the given dataset (all elements should maintain this criteria [root element is always smaller than its children])
2. remove the first element from the heap, put the last element at top and reheapify the heap
3. heapify starts from inner most node that is (size/2 - 1)
*/
