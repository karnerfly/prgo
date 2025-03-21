package sort

// Classic Quick Sort Function. Time Complexity - O(nlog(n)), Space Complexity - O(1)
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

/*  How Implemented?

1. take the pivot from the dataset
2. take two pointers start and end at low and high travers from both the side untils start <= end
3. where this criteria breaks (left element is less than the pivot and right element is greater than the pivot), swap the start and end (if start <= end) and push them forward
4. again do this process for sub dataset from low to end and from start to high until one element is remains in the sub dataset
*/
