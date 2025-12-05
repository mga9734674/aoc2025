package puzzle3

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strings"
)

func ParseAndRun(path string, size int) (int, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), "\n")

	banks := make([][]int, 0, len(splitted))

	for _, r := range splitted {
		n := len(r)
		if n < size {
			return 0, errors.New("invalid bank length")
		}

		bank := make([]int, n)

		for i := range n {
			bank[i] = int(r[i] - '0')
		}

		banks = append(banks, bank)
	}

	res := Run(banks, size)

	return res, nil
}

func Run(banks [][]int, size int) int {
	s := 0

	for _, bank := range banks {
		s += largestNumberInBank(bank, size)
		fmt.Println(s, bank)
	}

	return s
}

func largestNumberInBank(b []int, size int) int {
	n := len(b)
	m := make([]int, size)

	for i := range size {
		m[size-i-1] = b[n-i-1]
	}

	for i := n - size - 1; i >= 0; i-- {
		x := b[i]
		lastWasChanged := false

		for j := range size {
			if x >= m[j] && (lastWasChanged || j == 0) {
				t := m[j]
				m[j] = x
				x = t
				lastWasChanged = true
			} else {
				lastWasChanged = false
			}
		}
	}

	s := 0

	for i := range size {
		s += int(math.Pow10(size-i-1)) * m[i]
	}

	return s
}
