package _0587_erect_the_fence

import "sort"

type Point struct {
	X int
	Y int
}

// CrossProduct Determine the Cross Product of AB and BC.
// The sign of the cross product between AB and BC indicates orientation
//
// Returns an Orientation which indicates either:
// 	- AB x BC = 0 => CoLinear,
// 	- AB x BC < 0 => Clockwise
// 	- AB x BC > 0 =>  CounterClockwise
func CrossProduct(a, b, c Point) int {
	v1x, v1y := b.X-a.X, b.Y-a.Y
	v2x, v2y := c.X-b.X, c.Y-b.Y
	return v1x*v2y - v1y*v2x
}

func ConvexHull(points []Point) []Point {
	ans := make([]Point, 0)
	// Sort points in lexicographical order
	sort.Slice(points, func(i, j int) bool {
		return points[i].X < points[j].X || (points[i].X == points[j].X && points[i].Y < points[j].Y)
	})
	// Compute the Lower Hull, w/ ascending X
	for _, p := range points {
		// w/ enough elements, pop while AB BC forms a clockwise turn
		for len(ans) > 1 && CrossProduct(ans[len(ans)-2], ans[len(ans)-1], p) < 0 {
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, p)
	}
	if len(ans) == len(points) {
		return ans // All points are collinear and are included in hull
	}
	// Compute The Upper Hull w/ descending x, Avoid duplicating last point w/ -2
	for i := len(points) - 2; i >= 0; i-- {
		p := points[i]
		// w/ enough elements, pop while AB BC forms a clockwise turn
		for len(ans) > 1 && CrossProduct(ans[len(ans)-2], ans[len(ans)-1], p) < 0 {
			ans = ans[:len(ans)-1]
		}
		ans = append(ans, p)
	}
	ans = ans[:len(ans)-1] // Pop duplicate of first point
	sort.Slice(ans, func(i, j int) bool {
		return ans[i].X < ans[j].X || (ans[i].X == ans[j].X && ans[i].Y < ans[j].Y)
	})
	return ans
}

func MakeSlices(points []Point) [][]int {
	slices := make([][]int, len(points))
	for i, p := range points {
		slices[i] = []int{p.X, p.Y}
	}
	return slices
}

func MakePoints(trees [][]int) []Point {
	points := make([]Point, len(trees))
	for i, t := range trees {
		points[i] = Point{t[0], t[1]}
	}
	return points
}

func outerTrees(trees [][]int) [][]int {
	if len(trees) < 3 {
		return trees
	}
	points := MakePoints(trees)
	hull := ConvexHull(points)
	return MakeSlices(hull)
}
