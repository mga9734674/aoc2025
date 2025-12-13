package puzzle9

import (
	"bytes"
	"fmt"
	"math"
)

type Point struct {
	X, Y int
}

func Run(x []byte) (int, error) {
	rows := bytes.SplitSeq(bytes.Trim(x, "\n"), []byte{10})
	points := make([]Point, 0)

	for row := range rows {
		nums := bytes.Split(row, []byte{44})
		point := Point{X: toInt(nums[0]), Y: toInt(nums[1])}
		points = append(points, point)
	}

	b1, b2 := extremes(points)

	p1, p2 := outsiders(points, b1)
	p3, p4 := outsiders(points, b2)

	fmt.Println(p1, p2)
	fmt.Println(p3, p4)

	return max(area(p1, p2), area(p3, p4)), nil
}

func outsiders(points []Point, point Point) (Point, Point) {
	mini, maxi := math.MaxInt, 0
	idxMin, idxMax := 0, 0
	n := len(points)

	for i := range n {
		d := dist(points[i], point)

		if d < mini {
			idxMin = i
			mini = d
		}

		if d > maxi {
			idxMax = i
			maxi = d
		}
	}

	return points[idxMin], points[idxMax]
}

func extremes(points []Point) (Point, Point) {
	xMin, xMax, yMin, yMax, n := math.MaxInt, 0, math.MaxInt, 0, len(points)

	for i := range n {
		if points[i].X < xMin {
			xMin = points[i].X
		}

		if points[i].X > xMax {
			xMax = points[i].X
		}

		if points[i].Y < yMin {
			yMin = points[i].Y
		}

		if points[i].Y > yMax {
			yMax = points[i].Y
		}
	}

	return Point{X: xMin, Y: yMin}, Point{X: xMax, Y: yMin}
}

func area(p, q Point) int {
	a := (p.X - q.X)
	b := (p.Y - q.Y)

	if a < 0 {
		a *= -1
	}

	if b < 0 {
		b *= -1
	}

	return (a + 1) * (b + 1)
}

func dist(p, q Point) int {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y)
}

func toInt(x []byte) int {
	s, n := 0, len(x)

	for i := range n {
		s += int(x[i]-'0') * int(math.Pow10(n-i-1))
	}

	return s
}
