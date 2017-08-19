package main

import "fmt"
import "golang.org/x/tour/tree"
import "strconv"

// Simple point type, with offset operation
type point struct {
	x int
	y int
}

func (p point) offset(x int, y int) point {
	return point{p.x + x, p.y + y}
}

// Helper functions

// intDigits returns digits of an integer written in decimal
func intDigits(a int) (digits int) {
	digits = len([]rune(strconv.Itoa(a)))
	return
}

// max returns of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// min returns min of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// treeMaxValue returns the max value of a tree
func treeMaxValue(root *tree.Tree) int {
	if root == nil {
		return -1001
	}
	if isLeaf(root) {
		return root.Value
	}
	ret := root.Value
	maxLeftRight := max(treeMaxValue(root.Left), treeMaxValue(root.Right))
	ret = max(ret, maxLeftRight)
	return ret
}

// isLeaf returns true if root.Left == root.Right == nil. False otherwise
func isLeaf(root *tree.Tree) (ret bool) {
	return root.Left == nil && root.Right == nil
}

// Main functions

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
func maxLeft(root *tree.Tree) int {
	if root == nil {
		return 0
	}
	if isLeaf(root) {
		return 1
	}
	ret := 1 + max(maxLeft(root.Left), maxLeft(root.Right)-1)
	return ret
}

// maxRight returns max right offset. See maxLeft.
func maxRight(root *tree.Tree) (ret int) {
	if root == nil {
		return 0
	}
	if isLeaf(root) {
		ret = 1
		return
	}
	ret = 1 + max(maxRight(root.Left)-2, maxRight(root.Right))
	return
}

// Map containing runes at specific points.
// More flexible than a 2D array or 2D slices,
// to be converted into a string
var points map[point]rune

// PrettyTree returns a multi-line string representation of the tree
func PrettyTree(root *tree.Tree) string {
	points = make(map[point]rune) // Clear points map

	addTreePoints(point{0, 0}, root) // Add points
	xMin := 0
	xMax := 0
	yMax := 0

	for p := range points {
		yMax = max(yMax, p.y)
		xMin = min(xMin, p.x)
		xMax = max(xMax, p.x)
	}
	fmt.Println("yMax:", yMax)
	fmt.Println("xMax:", xMax)
	fmt.Println("xMin:", xMin)

	rect := make([][]rune, yMax+1)
	for s := range rect {
		rect[s] = make([]rune, xMax-xMin+1)
		for i := range rect[s] {
			rect[s][i] = ' '
		}
	}

	for p, v := range points {
		rect[p.y][p.x-xMin] = v
	}

	ret := ""
	for i := range rect {
		ret = ret + string(rect[i]) + "\n"
	}

	return ret
}

// addString adds necessary runes at correct points to represent string
func addString(p point, value string) {
	offset := (len(value) - 1) / 2
	for index, char := range value {
		points[p.offset(-offset+index, 0)] = char
	}
}

// addLines adds `count` left and right lines starting at point `p`
func addLines(p point, count int) {
	for i := 1; i <= count; i++ {
		points[p.offset(-i, i)] = '/'
		points[p.offset(i, i)] = '\\'
	}
}

// addTreePoints is the main helper function. Recursively adds tree node points
// at the correct positions, as well as lines connecting them
func addTreePoints(p point, node *tree.Tree) {
	if node == nil {
		return
	}

	addString(p, strconv.Itoa(node.Value))

	if isLeaf(node) {
		return
	}

	if node.Left == nil { // node.Right != nil
		points[p.offset(1, 1)] = '\\'
		addTreePoints(p.offset(2, 2), node.Right)
	} else if node.Right == nil { // node.Left != nil
		points[p.offset(-1, 1)] = '/'
		addTreePoints(p.offset(-2, 2), node.Left)
	} else { // both sides exist
		lines := max(1, maxRight(node.Left)+maxLeft(node.Right)-2)
		k := intDigits(treeMaxValue(node))
		fmt.Println("Value of k*lines:", k*lines)
		offset := k * lines
		addLines(p, offset)
		addTreePoints(p.offset(-offset-1, offset+1), node.Left)
		addTreePoints(p.offset(offset+1, offset+1), node.Right)
	}
}

func main() {
	root := tree.New(100)

	fmt.Println(PrettyTree(root))
}