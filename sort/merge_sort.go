package sort

// Classic Inplace Merge Sort Function. Time Complexity - O(nlog(n)), Space Complexity - O(log(n))
func MergeSort(data []int) {
	start := 0
	end := len(data) - 1
	mergeSort(data, start, end)
}

func mergeSort(arr []int, start, end int) {
	if start >= end {
		return
	}

	mid := start + (end-start)/2
	mergeSort(arr, start, mid)
	mergeSort(arr, mid+1, end)
	merge(arr, start, mid, mid+1, end)
}

func merge(arr []int, l1, u1, l2, u2 int) {
	tmp := make([]int, u2-l1+1)

	i, j, k := l1, l2, 0
	for i <= u1 && j <= u2 {
		if arr[i] <= arr[j] {
			tmp[k] = arr[i]
			k++
			i++
		} else {
			tmp[k] = arr[j]
			k++
			j++
		}
	}

	for i <= u1 {
		tmp[k] = arr[i]
		k++
		i++
	}

	for j <= u2 {
		tmp[k] = arr[j]
		k++
		j++
	}

	for i, k = l1, 0; i <= u2; i, k = i+1, k+1 {
		arr[i] = tmp[k]
	}
}
