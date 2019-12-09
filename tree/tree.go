package tree

import (
	"errors"
)

type IntBinaryTree struct {
	Value int
	Left  *IntBinaryTree
	Right *IntBinaryTree
}

type BTree interface {
	// Search search a key in tree
	Search(interface{}) (interface{}, error)
	Insert(interface{}, interface{})
	NewNode() *BTreeNode
	Split(*BTreeNode, interface{})
	Delete()
}

const (
	BTREE_MIN_DEGREE = 2
)

const (
	BTREE_ERROR_INVALID_DEGREE = errors.New("invalid degreee")
)
