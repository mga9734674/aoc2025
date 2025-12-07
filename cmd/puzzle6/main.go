package main

import (
	"fmt"
	"os"

	"aoc2025/internal/puzzle6"
)

const inputPath = `/tmp/aoc/puzzle6`

func main() {
	x, err := os.ReadFile(inputPath)
	if err != nil {
		panic(err)
	}

	res, err := puzzle6.ParseAndRun(x)
	if err != nil {
		panic(err)
	}

	fmt.Println("result", res)
}
