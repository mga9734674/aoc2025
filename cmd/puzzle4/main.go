package main

import (
	"fmt"

	"aoc2025/internal/puzzle4"
)

const inputPath = `/tmp/aoc/puzzle4`

func main() {
	res, err := puzzle4.ParseAndRun(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
