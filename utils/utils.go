package utils

import "strconv"
import "github.com/teolandon/trees/trees"

// intDigits returns digits of an integer written in decimal
func IntDigits(a int) (digits int) {
	digits = len([]rune(strconv.Itoa(a)))
	return
}

// Max returns of two integers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns min of two integers
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// treeMaxValue returns the max value of a tree
func TreeMaxValueLength(root *trees.Tree) int {
	if root == nil {
		return 0
	}
	ret := len(root.Value.String())
	maxLeftRight := Max(TreeMaxValueLength(root.Left), TreeMaxValueLength(root.Right))
	ret = Max(ret, maxLeftRight)
	return ret
}

// maxLeft returns how "left" the leftmost node in Tree root is.
// Meaning, how many units to the left of the root node
// is the leftmost node in a graphical representation
// of the tree located.
//
// More formally, it returns the maximum "left offset" value
// of all the nodes in the tree, where "left offset" is defined
// as the amount of left movements minus the amount of right movements
// that are required to reach a node, starting from the root node.
// The root node is counted as a movement in itself, so a leaf will
// return 1, not 0.
//
// I'm actually pretty bad at explaining this, but it's not exported
// anyway, so I'll leave it for later.
func MaxLeft(root *trees.Tree) int {
	if root == nil {
		return 0
	}
	if root.IsLeaf() {
		return 1
	}
	ret := 1 + Max(MaxLeft(root.Left), MaxLeft(root.Right)-1)
	return ret
}

// MaxRight returns max right offset. See MaxLeft.
func MaxRight(root *trees.Tree) (ret int) {
	if root == nil {
		return 0
	}
	if root.IsLeaf() {
		ret = 1
		return
	}
	ret = 1 + Max(MaxRight(root.Left)-2, MaxRight(root.Right))
	return
}
