package puzzle1

import (
	"errors"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const PointIni = 50

type Rotation struct {
	IsRight bool
	N       int
}

func ParseAndRun(path string) (int, error) {
	payload, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}

	splitted := strings.Split(strings.TrimRight(string(payload), "\n"), "\n")

	fmt.Println(splitted)

	rotations := make([]Rotation, 0, len(splitted))

	for _, seq := range splitted {
		rotation := Rotation{}

		if seq[0:1] == "R" {
			rotation.IsRight = true
		}

		n, err := strconv.Atoi(string(seq[1:]))
		if err != nil {
			return 0, errors.New("could not parse seq number")
		}

		rotation.N = n

		rotations = append(rotations, rotation)
	}

	res := Run(rotations)

	fmt.Println("n", len(rotations))

	return res, nil
}

func Run(rotations []Rotation) int {
	s := 0

	currPoint := PointIni

	for _, r := range rotations {
		if r.IsRight {
			currPoint = (currPoint + r.N) % 100
		} else {
			currPoint = int(math.Abs(float64((100 - (r.N%100 - currPoint))))) % 100
		}

		if currPoint == 0 {
			s++
		}

		fmt.Println(currPoint)
	}

	return s
}
