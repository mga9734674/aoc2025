package main

import (
	"fmt"

	"aoc2025/internal/puzzle2"
)

const inputPath = `/tmp/aoc/puzzle2`

func main() {
	res, err := puzzle2.ParseAndRun(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
