package trees

import "github.com/teolandon/trees/comparator"

type Tree struct {
	Left       *Tree
	Value      interface{}
	Right      *Tree
	comparator comparator.Comparator
}

// IsLeaf returns true if root.Left == root.Right == nil. False otherwise
// if root = nil, it returns false.
func (root *Tree) IsLeaf() (ret bool) {
	if root == nil {
		return false
	}
	return root.Left == nil && root.Right == nil
}
