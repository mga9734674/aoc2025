package main

import (
	"fmt"

	"aoc2025/internal/puzzle1"
)

const inputPath = `/tmp/aoc/puzzle1`

func main() {
	res, err := puzzle1.ParseAndRun(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
