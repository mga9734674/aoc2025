package puzzle3

import (
	"errors"
	"os"
	"strings"
)

func ParseAndRun(path string) (int, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), "\n")

	banks := make([][]int, 0, len(splitted))

	for _, r := range splitted {
		n := len(r)
		if n < 2 {
			return 0, errors.New("invalid bank length")
		}

		bank := make([]int, n)

		for i := range n {
			bank[i] = int(r[i] - '0')
		}

		banks = append(banks, bank)
	}

	res := Run(banks)

	return res, nil
}

func Run(banks [][]int) int {
	s := 0

	for _, bank := range banks {
		s += largestNumberInBank(bank)
	}

	return s
}

func largestNumberInBank(b []int) int {
	n := len(b)
	m1, m2 := b[n-2], b[n-1]

	for i := n - 3; i >= 0; i-- {
		if b[i] >= m1 {
			t := m1

			m1 = b[i]
			if t > m2 {
				m2 = t
			}
		}
	}

	return 10*m1 + m2
}
