package sort

// Classic Selection Sort Function. Time Complexity - O(n^2), Space Complexity - O(1)
func SelectionSort(data []int) {
	var (
		min  int
		size = len(data)
	)
	for i := 0; i < size; i++ {
		min = findMin(data, i, size-1)
		swap(data, i, min)
	}
}

func findMin(arr []int, i, n int) int {
	min := i
	for j := i + 1; j <= n; j++ {
		if arr[j] < arr[min] {
			min = j
		}
	}
	return min
}
