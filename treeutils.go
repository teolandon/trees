package main

import "fmt"
import "golang.org/x/tour/tree"

type Point struct {
	x int
	y int
}

func (p Point) offset(x int, y int) Point {
	return Point{p.x + x, p.y + y}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func isLeaf(root *tree.Tree) (ret bool) {
	return root.Left == nil && root.Right == nil
}

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

var points = make(map[Point]rune)

func fmtTree(root *tree.Tree) string {
	printHelper(Point{0, 0}, root)
	xMin := 0
	xMax := 0
	yMax := 0
	for p := range points {
		yMax = max(yMax, p.y)
		xMin = min(xMin, p.x)
		xMax = max(xMax, p.x)
	}
	rect := make([][]rune, yMax+1)
	for s := range rect {
		rect[s] = make([]rune, xMax-xMin+1)
		for i := range rect[s] {
			rect[s][i] = ' '
		}
	}
	for p, v := range points {
		fmt.Println("Inserting point", p, "with value", v)
		rect[p.y][p.x-xMin] = v
	}
	ret := ""
	for i := range rect {
		ret = ret + string(rect[i]) + "\n"
	}

	return ret
}

func printHelper(p Point, node *tree.Tree) {
	if node == nil {
		return
	}
	points[p] = rune(node.Value + '0')
	if isLeaf(node) {
		return
	}
	if node.Left == nil { // node.Right != nil
		points[p.offset(1, 1)] = '\\'
		printHelper(p.offset(2, 2), node.Right)
	} else if node.Right == nil { // node.Left != nil
		points[p.offset(-2, 2)] = '/'
		printHelper(p.offset(-2, 2), node.Left)
	} else { // both sides exist
		lines := max(1, maxRight(node.Left)+maxLeft(node.Right)-2)
		addLines(p, lines)
		printHelper(p.offset(-lines-1, lines+1), node.Left)
		printHelper(p.offset(lines+1, lines+1), node.Right)
	}
}

func addLines(p Point, count int) {
	for i := 1; i <= count; i++ {
		points[p.offset(-i, i)] = '/'
		points[p.offset(i, i)] = '\\'
	}
}

func main() {
	root := tree.New(3)

	fmt.Println(fmtTree(root))
}
