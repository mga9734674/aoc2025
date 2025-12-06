package puzzle5

import (
	"errors"
	"os"
	"sort"
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

	res := Run(ranges)

	return res, nil
}

func Run(ranges []Range) int {
	// [3 5]
	// [4 9]
	// [12 13]
	// [11 14]
	// [4 15]
	// [1 14]

	// seen = [[3 5]]
	// seen = [[3 9]]
	// seen = [[3 9] [12 13]]
	// seen = [[3 9] [11 14]]
	// seen = [[3 15] [11 14]] -> [[3 15]]
	// seen = [[1 15]]

	seen := make([]Range, 0) // seen merged sorted
	for _, r := range ranges {
		seen = add(seen, r.Start, r.End)
	}

	s := 0

	for _, r := range seen {
		s += r.End - r.Start + 1
	}

	return s
}

func add(seen []Range, start, end int) []Range {
	var res []Range

	for _, r := range seen {
		if start > r.End {
			res = append(res, r)

			continue
		}

		if end < r.Start {
			res = append(res, Range{start, end})
			start, end = r.Start, r.End

			continue
		}

		// merge
		if start > r.Start {
			start = r.Start
		}

		if end < r.End {
			end = r.End
		}
	}

	res = append(res, Range{start, end})
	sort.Slice(res, func(i, j int) bool { return res[i].Start < res[j].Start })

	return res
}
