package puzzle8

import (
	"bytes"
	"maps"
	"math"
	"sort"
)

type Point struct {
	X, Y, Z float64
}

type Dist struct {
	D    float64
	I, J int
}

func Run(x []byte) (int, error) {
	rows := bytes.SplitSeq(bytes.Trim(x, "\n"), []byte{10})
	points := make([]Point, 0)

	for row := range rows {
		nums := bytes.Split(row, []byte{44})
		points = append(
			points,
			Point{X: toFloat64(nums[0]), Y: toFloat64(nums[1]), Z: toFloat64(nums[2])},
		)
	}

	n := len(points)
	d := make([]Dist, 0)

	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			d = append(d, Dist{D: dist(points[i], points[j]), I: i, J: j})
		}
	}

	sort.Slice(d, func(i, j int) bool { return d[i].D < d[j].D })

	circuits := make(map[int]map[int]struct{})

	for i := range n {
		addAndMerge(circuits, d[i].I, d[i].J)
	}

	lens := make(map[int]struct{})

	for _, v := range circuits {
		lens[len(v)] = struct{}{}
	}

	p := 1

	for range 3 {
		p *= getMaxAndDelete(lens)
	}

	return p, nil
}

func getMaxAndDelete(a map[int]struct{}) int {
	if len(a) == 0 {
		return 1
	}

	m := 0

	for k := range a {
		if k > m {
			m = k
		}
	}

	delete(a, m)

	return m
}

func addAndMerge(m map[int]map[int]struct{}, x, y int) {
	if _, ok := m[x]; !ok {
		m[x] = map[int]struct{}{x: {}, y: {}}
	}
	if _, ok := m[y]; !ok {
		m[y] = map[int]struct{}{x: {}, y: {}}
	}

	m3 := make(map[int]struct{})

	maps.Copy(m3, m[x])
	maps.Copy(m3, m[y])

	for k := range m3 {
		m[k] = m3
	}
}

func toFloat64(x []byte) float64 {
	s, n := 0.0, len(x)

	for i := range n {
		s += float64(x[i]-'0') * math.Pow10(n-i-1)
	}

	return s
}

func dist(p, q Point) float64 {
	return (p.X-q.X)*(p.X-q.X) + (p.Y-q.Y)*(p.Y-q.Y) + (p.Z-q.Z)*(p.Z-q.Z)
}
