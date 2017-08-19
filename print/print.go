package print

import "fmt"
import "github.com/teolandon/trees/trees"
import "github.com/teolandon/trees/utils"

// Simple point type, with offset operation
type point struct {
	x int
	y int
}

func (p point) offset(x int, y int) point {
	return point{p.x + x, p.y + y}
}

// Map containing runes at specific points.
// More flexible than a 2D array or 2D slices,
// to be converted into a string
var points map[point]rune

// PrettyTree returns a multi-line string representation of the tree
func PrettyTree(root *trees.Tree) string {
	points = make(map[point]rune) // Clear points map

	addTreePoints(point{0, 0}, root) // Add points
	xMin := 0
	xMax := 0
	yMax := 0

	for p := range points {
		yMax = utils.Max(yMax, p.y)
		xMin = utils.Min(xMin, p.x)
		xMax = utils.Max(xMax, p.x)
	}

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
func addTreePoints(p point, node *trees.Tree) {
	if node == nil {
		return
	}

	addString(p, node.Value.String())

	if node.IsLeaf() {
		return
	}

	if node.Left == nil { // node.Right != nil
		points[p.offset(1, 1)] = '\\'
		addTreePoints(p.offset(2, 2), node.Right)
	} else if node.Right == nil { // node.Left != nil
		points[p.offset(-1, 1)] = '/'
		addTreePoints(p.offset(-2, 2), node.Left)
	} else { // both sides exist
		lines := utils.Max(1, utils.MaxRight(node.Left)+utils.MaxLeft(node.Right)-2)
		k := utils.IntDigits(utils.TreeMaxValueLength(node))
		fmt.Println("Value of k*lines:", k*lines)
		offset := k * lines
		addLines(p, offset)
		addTreePoints(p.offset(-offset-1, offset+1), node.Left)
		addTreePoints(p.offset(offset+1, offset+1), node.Right)
	}
}
