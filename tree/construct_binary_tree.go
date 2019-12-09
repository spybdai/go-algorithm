package tree

import ()

// ConstructBinaryTreeFromPreAndInWrap is a wrapper
func ConstructBinaryTreeFromPreAndInWrap(pre, in []int) *IntBinaryTree {
	index, start := 0, 0
	end := len(in) - 1

	tree := ConstructBinaryTreeFromPreAndIn(pre, in, &index, start, end)
	return tree
}

// ConstructBinaryTreeFromPreAndIn construct binary tree from pre-order and in-order array
// index is the index of pre-order array
// start and end are the pointers point to start and index of in-order array
// time complexity around O(nlogn) (for any note, deepth of sub-left and sub-right no more than 1)
// can downgrade to O(n**2)
func ConstructBinaryTreeFromPreAndIn(pre, in []int, index *int, start, end int) *IntBinaryTree {

	node := &IntBinaryTree{
		Value: pre[*index],
	}

	if start == end {
		return node
	}

	rootIndex := SearchIntArray(in, pre[*index], 0, len(in)-1)
	if rootIndex == -1 {
		return nil
	}

	if start == rootIndex {
		node.Left = nil
	} else {
		*index++
		node.Left = ConstructBinaryTreeFromPreAndIn(pre, in, index, start, rootIndex-1)
	}

	if rootIndex == end {
		node.Right = nil
	} else {
		*index++
		node.Right = ConstructBinaryTreeFromPreAndIn(pre, in, index, rootIndex+1, end)
	}

	return node
}

// SearchIntArray search the index of given value from start to end
func SearchIntArray(array []int, value, start, end int) int {
	for index := start; index <= end; index++ {
		if value == array[index] {
			return index
		}
	}
	return -1
}
