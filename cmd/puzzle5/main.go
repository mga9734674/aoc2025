package main

import (
	"fmt"

	"aoc2025/internal/puzzle5"
)

const inputPath = `/tmp/aoc/puzzle5`

func main() {
	res, err := puzzle5.ParseAndRun(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
