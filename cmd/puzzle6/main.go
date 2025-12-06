package main

import (
	"fmt"

	"aoc2025/internal/puzzle6"
)

const inputPath = `/tmp/aoc/puzzle6`

func main() {
	res, err := puzzle6.ParseAndRun(inputPath)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
