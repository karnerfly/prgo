package sort

func QuickSort(data []int) {
	low := 0
	high := len(data) - 1
	quickSort(data, low, high)
}

func quickSort(arr []int, low, high int) {

	if low == high {
		return
	}

	start := low
	end := high
	mid := start + (end-start)/2
	pivot := arr[mid]

	for start <= end {
		for arr[start] < pivot {
			start++
		}

		for arr[end] > pivot {
			end--
		}

		if start <= end {
			swap(arr, start, end)
			start++
			end--
		}
	}

	quickSort(arr, low, end)
	quickSort(arr, start, high)
}
