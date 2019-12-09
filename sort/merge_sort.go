package sort

import ()

// MergeSort
// Recursive Approach
// tc: O(nlgn)
// key point: slice edge
func MergeSort(array []int, ws []int, start, end int) {

	mid := (end-start)>>1 + start
	if start >= end {
		return
	}

	MergeSort(array, ws, start, mid)
	MergeSort(array, ws, mid+1, end)

	i, j, index := start, mid+1, start

	for ; i <= mid && j <= end; index++ {
		if array[i] <= array[j] {
			ws[index] = array[i]
			i++
		} else {
			ws[index] = array[j]
			j++
		}
	}

	for ; i <= mid; i++ {
		ws[index] = array[i]
		index++
	}

	for ; j <= end; j++ {
		ws[index] = array[j]
		index++
	}

	for index = start; index <= end; index++ {
		array[index] = ws[index]
	}
}
