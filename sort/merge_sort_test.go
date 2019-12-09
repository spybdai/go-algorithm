package sort

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	array := []int{7, 3, 1, 5, 4, 0, 9, 6, 8, 2}
	ws := make([]int, len(array))

	MergeSort(array, ws, 0, len(array)-1)
	for index := 0; index < len(array)-1; index++ {
		assert.Equal(t, true, array[index] <= array[index+1])
	}

}
