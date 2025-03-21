package sort

// Classic Insertion Sort Function. Time Complexity - O(n^2), Space Complexity - O(1)
func InsertionSort(data []int) {
	var (
		j    int
		key  int
		size = len(data)
	)
	for i := 1; i < size; i++ {
		j = i - 1
		key = data[i]
		insert(data, j, key)
	}
}

func insert(data []int, n, key int) {
	var i int
	for i = n; i >= 0; i-- {
		if data[i] < key {
			break
		}
		data[i+1] = data[i]
	}
	data[i+1] = key
}
