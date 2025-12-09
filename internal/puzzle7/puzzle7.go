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

	s := 0
	beams := []int{p / 2}

	for i := 2 * p; i < n*p; i += p {
		beamsNew := make([]int, 0)
		seen := make(map[int]struct{})

		for _, b := range beams {
			j := i + b

			if x[j] == '^' {
				s++

				if !has(seen, b-1) {
					beamsNew = append(beamsNew, b-1)
					seen[b-1] = struct{}{}
				}

				beamsNew = append(beamsNew, b+1)
			}

			if x[j] == '.' && !has(seen, b) {
				beamsNew = append(beamsNew, b)
				seen[b] = struct{}{}
			}
		}

		beams = beamsNew
	}

	return s, nil
}

func has(m map[int]struct{}, i int) bool {
	if _, ok := m[i]; ok {
		return true
	}

	return false
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
