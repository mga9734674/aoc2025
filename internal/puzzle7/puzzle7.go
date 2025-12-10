package puzzle7

import (
	"bytes"
	"fmt"
)

func Run(x []byte) (int, error) {
	n := countNl(x)
	x = bytes.ReplaceAll(x, []byte{10}, nil)
	p := len(x) / n

	if len(x)%n != 0 {
		return 0, fmt.Errorf("invalid format %d %d %d", len(x), n, len(x)%n)
	}

	seen := make(map[int]int)

	return run(x, seen, n, p, 1, p/2), nil
}

func run(x []byte, seen map[int]int, n, p, i, j int) int {
	if i < 0 || i > n-1 || j < 0 || j > p-1 {
		return 1
	}

	idx := i*p + j

	if v, ok := seen[i*p+j]; ok {
		return v
	}

	s := 0

	if x[idx] == '^' {
		s += run(x, seen, n, p, i+1, j-1) + run(x, seen, n, p, i+1, j+1)
	} else {
		s += run(x, seen, n, p, i+1, j)
	}

	seen[idx] = s

	return s
}

func countNl(v []byte) int {
	s := 0

	for _, x := range v {
		if x == 10 {
			s++
		}
	}

	return s
}
