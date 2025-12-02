package puzzle2

import (
	"errors"
	"fmt"
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

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), ",")

	ranges := make([]Range, 0, len(splitted))

	for _, r := range splitted {
		x := strings.Split(r, "-")
		if len(x) != 2 {
			return 0, errors.New("invalid range")
		}

		start, err := strconv.Atoi(x[0])
		if err != nil {
			return 0, err
		}

		end, err := strconv.Atoi(x[1])
		if err != nil {
			return 0, err
		}

		ranges = append(ranges, Range{Start: start, End: end})
	}

	res := Run(ranges)

	return res, nil
}

func Run(ranges []Range) int {
	s := 0

	for _, r := range ranges {
		for i := r.Start; i < r.End+1; i++ {
			id := strconv.Itoa(i)
			if len(id)%2 == 1 {
				continue
			}

			if id[0:len(id)/2] == id[len(id)/2:] {
				s += i

				fmt.Println(id)
			}
		}
	}

	return s
}
