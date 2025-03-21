package sort

type bin struct {
	data int
	next *bin
}

func newBin(data int) *bin {
	return &bin{data: data}
}

// Classic Selection Sort Function. Time Complexity - O(d*n) [d = no of digits in the max number], Space Complexity - O(n + r) [r = no of digits of the number system]
func RadixSort(data []int, base int) {
	bins := make([]*bin, base)
	maxNum := findMaxNum(data)

	for weight := 1; maxNum/weight > 0; weight *= base {
		for j := 0; j < len(data); j++ {
			put(bins, getDigit(data[j], weight), data[j])
		}
		rearrange(bins, data)
	}
}

func put(bins []*bin, i, data int) {
	if bins[i] == nil {
		bin := newBin(data)
		bins[i] = bin
	} else {
		ptr := bins[i]
		for ptr.next != nil {
			ptr = ptr.next
		}
		bin := newBin(data)
		ptr.next = bin
	}
}

func rearrange(bins []*bin, arr []int) {
	j := 0
	for i := 0; i < len(bins); i++ {
		if bins[i] != nil {
			arr[j] = bins[i].data
			j++

			if bins[i].next != nil {
				ptr := bins[i].next
				for ptr != nil {
					arr[j] = ptr.data
					j++
					ptr = ptr.next
				}
			}
		}
		// clean bins for next pass
		bins[i] = nil
	}
}

func getDigit(n, i int) int {
	return (n / i) % 10
}

func findMaxNum(arr []int) int {
	max := 0
	for j := 1; j < len(arr); j++ {
		if arr[j] > arr[max] {
			max = j
		}
	}
	return arr[max]
}

/*  How Implemented?

1. create a bin list length is the no of digits in that number system
2. first find the largest number from the dataset (outer loop runs for n times where n = no of digits in largest number)
3. for every element in the dataset get the significant digit value for current pass
4. put the element into the corresponding bucket of the digit (bin[digit] = item), if the bin already has any previous value then linked the new item at the end of that bin slot (bin[digit] <- item1 <- item2 <- newitem)
5. traverse the whole bins list from first to end and collect the items and put them suquentially into the dataset
*/
