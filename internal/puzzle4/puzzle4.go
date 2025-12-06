package puzzle4

import (
	"os"
	"strings"
)

func ParseAndRun(path string) (int, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), "\n")

	rows := make([][]bool, 0, len(splitted))

	for _, r := range splitted {
		n := len(r)

		row := make([]bool, n)

		for i := range n {
			if r[i] == '@' {
				row[i] = true
			}
		}

		rows = append(rows, row)
	}

	res := Run(rows)

	return res, nil
}

func Run(rows [][]bool) int {
	s := 0
	n := len(rows)

	for i := range n {
		p := len(rows[i])

		for j := range p {
			s1 := 0

			if !rows[i][j] {
				continue
			}

			// 1 2 3
			// 8 . 4
			// 7 6 5

			if i > 0 && j > 0 && rows[i-1][j-1] {
				s1++
			}

			if i > 0 && rows[i-1][j] {
				s1++
			}

			if i > 0 && j < p-1 && rows[i-1][j+1] {
				s1++
			}

			if j < p-1 && rows[i][j+1] {
				s1++
			}

			if i < n-1 && j < p-1 && rows[i+1][j+1] {
				s1++
			}

			if i < n-1 && rows[i+1][j] {
				s1++
			}

			if i < n-1 && j > 0 && rows[i+1][j-1] {
				s1++
			}

			if j > 0 && rows[i][j-1] {
				s1++
			}

			if s1 < 4 {
				s++
			}
		}
	}

	return s
}
