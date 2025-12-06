package puzzle6

import (
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var reNum = regexp.MustCompile(`([0-9]+|\*|\+)\s*`)

type Compute struct {
	Nums []int
	Ope  string
}

func ParseAndRun(path string) (int, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), "\n")
	n := len(splitted)
	if n < 3 {
		return 0, errors.New("invalid format")
	}
	rows := make([][]string, 0, n)

	for i := range n {
		match := reNum.FindAllStringSubmatch(splitted[i], -1)
		row := make([]string, 0, len(match))

		if i > 0 && len(match) != len(rows[i-1]) {
			return 0, errors.New("invalid length")
		}

		for _, m := range match {
			row = append(row, m[1])
		}

		rows = append(rows, row)
	}

	p := len(rows[0])
	c := make([]Compute, p)

	for j := range p {
		nums := make([]int, n-1)
		for i := range n - 1 {
			num, err := strconv.Atoi(rows[i][j])
			if err != nil {
				return 0, err
			}

			nums[i] = num
		}

		c[j] = Compute{Nums: nums, Ope: rows[n-1][j]}
	}

	res := Run(c)

	return res, nil
}

func Run(c []Compute) int {
	s := 0

	for _, x := range c {
		if x.Ope == "*" {
			s += mul(x.Nums)
		} else {
			s += sum(x.Nums)
		}
	}

	return s
}

func sum(nums []int) int {
	s := 0

	for _, num := range nums {
		s += num
	}

	return s
}

func mul(nums []int) int {
	s := 1

	for _, num := range nums {
		s *= num
	}

	return s
}
