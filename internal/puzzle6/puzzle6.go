package puzzle6

import (
	"bytes"
	"errors"
	"fmt"
	"math"
)

type Compute struct {
	Nums []int
	Op   string
}

func ParseAndRun(x []byte) (int, error) {
	n := countNl(x)
	x = bytes.ReplaceAll(x, []byte{10}, nil)
	p := len(x) / n

	if len(x)%n != 0 {
		return 0, errors.New("invalid format")
	}

	sepIndex := []int{-1}

	for j := range p {
		isEmpty := true

		for i := range n {
			if x[j+p*i] != 32 {
				isEmpty = false

				break
			}
		}

		if isEmpty {
			sepIndex = append(sepIndex, j)
		}
	}

	sepIndex = append(sepIndex, p)

	q := len(sepIndex)
	c := make([]Compute, 0)

	for k := 1; k < q; k++ {
		block := make([][]int, 0)

		for j := sepIndex[k-1] + 1; j < sepIndex[k]; j++ {
			digits := make([]int, 0)

			for i := range n - 1 {
				y := x[j+p*i]

				if y != 32 {
					digits = append(digits, int(y-'0'))
				}
			}

			block = append(block, digits)
		}

		c = append(c, buildCompute(block, string(x[sepIndex[k-1]+1+p*(n-1)])))
	}

	fmt.Println(c)

	res := run(c)

	return res, nil
}

func buildCompute(block [][]int, op string) Compute {
	n := len(block)
	nums := make([]int, n)

	for i := range n {
		p := len(block[i])
		num := 0

		for j := range p {
			num += int(math.Pow10(p-j-1)) * block[i][j]
		}

		nums[i] = num
	}

	return Compute{Nums: nums, Op: op}
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

func run(c []Compute) int {
	s := 0

	for _, x := range c {
		if x.Op == "*" {
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
