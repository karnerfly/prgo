package search

// Classic Binary Search Function. Time Complexity - O(log(n)), Space Complexity - O(1)
func BinarySearch(data []int, key int) int {
	var (
		mid   int
		start = 0
		end   = len(data) - 1
	)

	for start <= end {
		mid = start + (end-start)/2
		if data[mid] == key {
			return mid
		} else if data[mid] > key {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return -1
}
