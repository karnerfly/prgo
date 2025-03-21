package sort

// Classic Cycle Sort Function. Time Complexity - O(n^2), Space Complexity - O(1)
func CycleSort(data []int) {
	var (
		isSwap bool
		size   = len(data)
	)
	for i := 0; i < size; i++ {
		isSwap = false
		for j := 0; j < size; j++ {
			if data[j] != j+1 {
				swap(data, j, data[j]-1)
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}
