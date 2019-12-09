package tree

import ()

func PreOrder(tree *IntBinaryTree, preArray *[]int) {
	if tree != nil {
		*preArray = append(*preArray, tree.Value)
		PreOrder(tree.Left, preArray)
		PreOrder(tree.Right, preArray)
	}
}
