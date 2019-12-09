package tree

import (
	//"github.com/stretchr/testify/assert"
	"fmt"
	"testing"
)

func TestConstructBinaryTreeFromPreAndIn(t *testing.T) {
	in := []int{4, 8, 2, 5, 1, 6, 3, 7}
	pre := []int{1, 2, 4, 8, 5, 3, 6, 7}
	tree := ConstructBinaryTreeFromPreAndInWrap(pre, in)
	array := &[]int{}
	PreOrder(tree, array)
	fmt.Println(pre)
	fmt.Println(*array)
}
