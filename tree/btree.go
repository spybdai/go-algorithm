package tree

type BTreeNode struct {
	Leaf     bool
	Number   int32
	Key      []interface{}
	Value    []interface{}
	Children []*BTree
}

type BTreeO struct {
	Root             *BTreeNode
	Degree           int32
	CompareFuncEqual func(interface{}, interface{}) bool
	CompareFuncELT   func(interface{}, interface{}) bool
	CompareFuncGT    func(interface{}, interface{}) bool
}

type BTreeCompareFunc func(interface{}, interface{}) bool

func NewBTree(degree int32, function BTreeCompareFunc) (*BTree, error) {
	if degree < BTREE_MIN_DEGREE {
		return nil, BTREE_ERROR_INVALID_DEGREE
	}

	root := &BTreeNode{
		Leaf:   true,
		Number: 0,
		Key:    make([]interface{}, degree),
		Value:  make([]interface{}, degree),
	}

	tree := &BTreeO{
		Root:             root,
		Degree:           degree,
		CompareFuncEqual: function,
	}

	return tree, nil
}

func (bt *BTreeO) Search(key interface{}) (val interface{}, err error) {
	if bt.Root == nil {
		return
	}
	return Root.searchNode(bt.Root, key)
}

func (bt *BTreeO) nodeIsFull(btn *BTreeNode) bool {
	return btn.Number == bt.Degree
}

// searchNode search the nodes of btree
func (bt *BTreeO) searchNode(btn *BTreeNode, key interface{}) (val interface{}, err error) {
	for i := 0; i < btn.Number; i++ {
		if bt.CompareFuncELT(key, btn.Key[i]) {
			break
		}
	}
	if i < btn.Number && bt.CompareFuncEqual(key, btn.Key[i]) { // hit key in current node
		return btn.Value[i], nil
	} else if btn.Leaf { // if leaf node, then, not found
		return
	} else { // search child
		// read-disk(btn.Children[i]), if node not in memory
		return bt.searchNode(btn.Children[i], key)
	}
}

func (bt *BTreeO) splitNode(btn *BTreeNode, index int32) {
}
