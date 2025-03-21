package search

// Classic Linear Search Function. Time Complexity - O(n), Space Complexity - O(1)
func LinearSearch(data []int, key int) int {
	for i := 0; i < len(data); i++ {
		if data[i] == key {
			return i
		}
	}
	return -1
}
