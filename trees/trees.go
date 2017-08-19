package trees

import "fmt"

type Tree struct {
	Left  *Tree
	Value fmt.Stringer
	Right *Tree
}

// IsLeaf returns true if root.Left == root.Right == nil. False otherwise
// if root = nil, it returns false.
func (root *Tree) IsLeaf() (ret bool) {
	if root == nil {
		return false
	}
	return root.Left == nil && root.Right == nil
}
