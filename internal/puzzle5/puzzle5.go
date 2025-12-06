package puzzle5

import (
	"errors"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

func ParseAndRun(path string) (int, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), "\n\n")
	if len(splitted) != 2 {
		return 0, errors.New("invalid format")
	}

	x := strings.Split(splitted[0], "\n")
	ranges := make([]Range, 0, len(x))

	for _, r := range x {
		x0 := strings.Split(r, "-")
		if len(x0) != 2 {
			return 0, errors.New("invalid range")
		}

		start, err := strconv.Atoi(x0[0])
		if err != nil {
			return 0, err
		}

		end, err := strconv.Atoi(x0[1])
		if err != nil {
			return 0, err
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	y := strings.Split(splitted[1], "\n")
	ingredients := make([]int, 0, len(y))

	for _, r := range y {
		ingredient, err := strconv.Atoi(r)
		if err != nil {
			return 0, err
		}

		ingredients = append(ingredients, ingredient)

	}

	res := Run(ranges, ingredients)

	return res, nil
}

func Run(ranges []Range, ingredients []int) int {
	s := 0

	for _, ingredient := range ingredients {
		for _, r := range ranges {
			if ingredient >= r.Start && ingredient <= r.End {
				s++
				break
			}
		}
	}

	return s
}
