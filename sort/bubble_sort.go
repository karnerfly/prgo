package sort

// Classic Bubble Sort Function. Time Complexity - O(n^2), Space Complexity - O(1)
func BubbleSort(data []int) {
	var isSwap bool
	for i := 0; i < len(data)-1; i++ {
		isSwap = false
		for j := 0; j < len(data)-i-1; j++ {
			if data[j] > data[j+1] {
				swap(data, j, j+1)
				isSwap = true
			}
		}
		if !isSwap {
			break
		}
	}
}
